package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"
	"unaryrpc/pb"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type User struct {
	ID   string
	Name string
	Age  int32
}

type UserService struct {
	pb.UnimplementedUserServer

	users map[string]User
	mu    sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]User),
	}
}

func (us *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.Username == "admin" && in.Password == "admin" {
		s, err := generateJWT(in.Username)
		if err != nil {
			return nil, err
		}

		return &pb.LoginResponse{
			Token: s,
		}, nil
	}

	return nil, errors.New("invalid username or password")
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      jwt.TimeFunc().Add(time.Minute * 15).Unix(),
	})

	return token.SignedString(jwtKey)
}

func (us *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user := User{
		ID:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}

	us.users[user.ID] = user

	return &pb.AddUserResponse{
		Id:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}

func Run() {
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		return
	}

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))
	pb.RegisterUserServer(s, NewUserService())
	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user, ok := us.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &pb.GetUserResponse{
		Id:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
