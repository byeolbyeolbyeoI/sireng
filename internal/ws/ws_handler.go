package ws

import (
	"fmt"
	"github.com/chaaaeeee/sireng/util"
	"net/http"

	"github.com/gorilla/websocket"
)

type WsHandler interface {
	CreateRoom(w http.ResponseWriter, r *http.Request)
	GetRooms(w http.ResponseWriter, r *http.Request)
	JoinRoom(w http.ResponseWriter, r *http.Request)
	GetClients(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	hub  *Hub
	util util.Util
}

func NewHandler(h *Hub, util util.Util) WsHandler {
	return &Handler{
		hub:  h,
		util: util,
	}
}

type CreateRoomReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomReq
	fmt.Println("here")

	if err := h.util.Input(r, &req); err != nil {
		h.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	h.hub.Rooms[req.Id] = &Room{
		Id:      req.Id,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	h.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Ok",
		Data:    req,
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tes")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	q := r.URL.Query()
	roomId := q.Get("roomId")
	clientId := q.Get("userId")
	username := q.Get("username")
	fmt.Println("roomid: ", roomId)
	fmt.Println("clientid: ", clientId)
	fmt.Println("username: ", username)

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		Id:       clientId,
		RoomId:   roomId,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomId:   roomId,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	h.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Successfully get rooms data",
		Data:    rooms,
	})
}

type ClientRes struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []ClientRes
	roomId := r.URL.Query().Get("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		h.util.WriteJSON(w, http.StatusOK, util.Response{
			Success: true,
			Message: "Successfully get clients data",
			Data:    clients,
		})
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			Id:       c.Id,
			Username: c.Username,
		})
	}

	h.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Successfully get clients data",
		Data:    clients,
	})
}
