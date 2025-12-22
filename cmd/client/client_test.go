package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	greetv1 "github.com/lao-tseu-is-alive/connect-go-example/gen/greet/v1"
	"github.com/lao-tseu-is-alive/connect-go-example/gen/greet/v1/greetv1connect"
)

// MockGreetServer is a simple in-memory implementation for testing
type MockGreetServer struct{}

func (s *MockGreetServer) Greet(ctx context.Context, req *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
	return &greetv1.GreetResponse{
		Greeting: "Hello, " + req.Name + "!",
	}, nil
}

// setupTestServer creates a test server and returns the client and cleanup function
func setupTestServer(t *testing.T, opts ...connect.ClientOption) (greetv1connect.GreetServiceClient, func()) {
	t.Helper()
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)

	server := httptest.NewServer(mux)

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
		opts...,
	)

	return client, server.Close
}

// --- Unit Tests ---

func TestGreet_Connect(t *testing.T) {
	client, cleanup := setupTestServer(t)
	defer cleanup()

	res, err := client.Greet(context.Background(), &greetv1.GreetRequest{Name: "TestUser"})
	if err != nil {
		t.Fatalf("Greet failed: %v", err)
	}

	expected := "Hello, TestUser!"
	if res.Greeting != expected {
		t.Errorf("Expected %q, got %q", expected, res.Greeting)
	}
}

func TestGreet_GRPC(t *testing.T) {
	client, cleanup := setupTestServer(t, connect.WithGRPC())
	defer cleanup()

	res, err := client.Greet(context.Background(), &greetv1.GreetRequest{Name: "GRPCUser"})
	if err != nil {
		t.Fatalf("Greet failed: %v", err)
	}

	expected := "Hello, GRPCUser!"
	if res.Greeting != expected {
		t.Errorf("Expected %q, got %q", expected, res.Greeting)
	}
}

func TestGreet_EmptyName(t *testing.T) {
	client, cleanup := setupTestServer(t)
	defer cleanup()

	res, err := client.Greet(context.Background(), &greetv1.GreetRequest{Name: ""})
	if err != nil {
		t.Fatalf("Greet failed: %v", err)
	}

	expected := "Hello, !"
	if res.Greeting != expected {
		t.Errorf("Expected %q, got %q", expected, res.Greeting)
	}
}

// --- Benchmark Tests ---

// BenchmarkGreet_Connect benchmarks the Connect protocol
func BenchmarkGreet_Connect(b *testing.B) {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	defer server.Close()

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Reset timer after setup
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Fatalf("Greet failed: %v", err)
		}
	}
}

// BenchmarkGreet_GRPC benchmarks the gRPC protocol
func BenchmarkGreet_GRPC(b *testing.B) {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	defer server.Close()

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
		connect.WithGRPC(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	// Reset timer after setup
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Fatalf("Greet failed: %v", err)
		}
	}
}

// BenchmarkGreet_Connect_Parallel benchmarks Connect with parallel requests
func BenchmarkGreet_Connect_Parallel(b *testing.B) {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	defer server.Close()

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Greet(context.Background(), req)
			if err != nil {
				b.Errorf("Greet failed: %v", err)
			}
		}
	})
}

// BenchmarkGreet_GRPC_Parallel benchmarks gRPC with parallel requests
func BenchmarkGreet_GRPC_Parallel(b *testing.B) {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	defer server.Close()

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
		connect.WithGRPC(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Greet(context.Background(), req)
			if err != nil {
				b.Errorf("Greet failed: %v", err)
			}
		}
	})
}

// --- JSON Benchmarks ---

// BenchmarkGreet_JSON benchmarks Connect with JSON encoding (sequential)
func BenchmarkGreet_JSON(b *testing.B) {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	defer server.Close()

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
		connect.WithProtoJSON(), // Use JSON encoding
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := client.Greet(context.Background(), req)
		if err != nil {
			b.Fatalf("Greet failed: %v", err)
		}
	}
}

// BenchmarkGreet_JSON_Parallel benchmarks Connect with JSON encoding (parallel)
func BenchmarkGreet_JSON_Parallel(b *testing.B) {
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&MockGreetServer{})
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	defer server.Close()

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		server.URL,
		connect.WithProtoJSON(),
	)

	req := &greetv1.GreetRequest{Name: "BenchUser"}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Greet(context.Background(), req)
			if err != nil {
				b.Errorf("Greet failed: %v", err)
			}
		}
	})
}
