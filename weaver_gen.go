// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(App{}),
		Listeners: []string{"listener"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦7cfe2a9e:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→hello/UserRepository⟧\n⟦2248fb79:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→listener⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "hello/UserRepository",
		Iface: reflect.TypeOf((*UserRepository)(nil)).Elem(),
		Impl:  reflect.TypeOf(userRepositoryImpl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return userRepository_local_stub{impl: impl.(UserRepository), tracer: tracer, createMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "Create", Remote: false}), deleteMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "Delete", Remote: false}), findAllMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "FindAll", Remote: false}), findByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "FindById", Remote: false}), updateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "Update", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return userRepository_client_stub{stub: stub, createMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "Create", Remote: true}), deleteMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "Delete", Remote: true}), findAllMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "FindAll", Remote: true}), findByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "FindById", Remote: true}), updateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/UserRepository", Method: "Update", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return userRepository_server_stub{impl: impl.(UserRepository), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return userRepository_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[weaver.Main] = (*App)(nil)
var _ weaver.InstanceOf[UserRepository] = (*userRepositoryImpl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*App)(nil)
var _ weaver.Unrouted = (*userRepositoryImpl)(nil)

// Local stub implementations.

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

type userRepository_local_stub struct {
	impl            UserRepository
	tracer          trace.Tracer
	createMetrics   *codegen.MethodMetrics
	deleteMetrics   *codegen.MethodMetrics
	findAllMetrics  *codegen.MethodMetrics
	findByIdMetrics *codegen.MethodMetrics
	updateMetrics   *codegen.MethodMetrics
}

// Check that userRepository_local_stub implements the UserRepository interface.
var _ UserRepository = (*userRepository_local_stub)(nil)

func (s userRepository_local_stub) Create(ctx context.Context, a0 User) (r0 User, err error) {
	// Update metrics.
	begin := s.createMetrics.Begin()
	defer func() { s.createMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.UserRepository.Create", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Create(ctx, a0)
}

func (s userRepository_local_stub) Delete(ctx context.Context, a0 uint64) (err error) {
	// Update metrics.
	begin := s.deleteMetrics.Begin()
	defer func() { s.deleteMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.UserRepository.Delete", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Delete(ctx, a0)
}

func (s userRepository_local_stub) FindAll(ctx context.Context) (r0 []User, err error) {
	// Update metrics.
	begin := s.findAllMetrics.Begin()
	defer func() { s.findAllMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.UserRepository.FindAll", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FindAll(ctx)
}

func (s userRepository_local_stub) FindById(ctx context.Context, a0 uint64) (r0 User, err error) {
	// Update metrics.
	begin := s.findByIdMetrics.Begin()
	defer func() { s.findByIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.UserRepository.FindById", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FindById(ctx, a0)
}

func (s userRepository_local_stub) Update(ctx context.Context, a0 User) (r0 User, err error) {
	// Update metrics.
	begin := s.updateMetrics.Begin()
	defer func() { s.updateMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.UserRepository.Update", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Update(ctx, a0)
}

// Client stub implementations.

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

type userRepository_client_stub struct {
	stub            codegen.Stub
	createMetrics   *codegen.MethodMetrics
	deleteMetrics   *codegen.MethodMetrics
	findAllMetrics  *codegen.MethodMetrics
	findByIdMetrics *codegen.MethodMetrics
	updateMetrics   *codegen.MethodMetrics
}

// Check that userRepository_client_stub implements the UserRepository interface.
var _ UserRepository = (*userRepository_client_stub)(nil)

func (s userRepository_client_stub) Create(ctx context.Context, a0 User) (r0 User, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.createMetrics.Begin()
	defer func() { s.createMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.UserRepository.Create", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_User_0df39765(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s userRepository_client_stub) Delete(ctx context.Context, a0 uint64) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.deleteMetrics.Begin()
	defer func() { s.deleteMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.UserRepository.Delete", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Uint64(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s userRepository_client_stub) FindAll(ctx context.Context) (r0 []User, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.findAllMetrics.Begin()
	defer func() { s.findAllMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.UserRepository.FindAll", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 2, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_User_42defd96(dec)
	err = dec.Error()
	return
}

func (s userRepository_client_stub) FindById(ctx context.Context, a0 uint64) (r0 User, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.findByIdMetrics.Begin()
	defer func() { s.findByIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.UserRepository.FindById", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Uint64(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s userRepository_client_stub) Update(ctx context.Context, a0 User) (r0 User, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.updateMetrics.Begin()
	defer func() { s.updateMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.UserRepository.Update", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_User_0df39765(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 4, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.22.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type userRepository_server_stub struct {
	impl    UserRepository
	addLoad func(key uint64, load float64)
}

// Check that userRepository_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*userRepository_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s userRepository_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Create":
		return s.create
	case "Delete":
		return s.delete
	case "FindAll":
		return s.findAll
	case "FindById":
		return s.findById
	case "Update":
		return s.update
	default:
		return nil
	}
}

func (s userRepository_server_stub) create(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 User
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Create(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s userRepository_server_stub) delete(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 uint64
	a0 = dec.Uint64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Delete(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s userRepository_server_stub) findAll(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.FindAll(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_User_42defd96(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s userRepository_server_stub) findById(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 uint64
	a0 = dec.Uint64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.FindById(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s userRepository_server_stub) update(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 User
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Update(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

type userRepository_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that userRepository_reflect_stub implements the UserRepository interface.
var _ UserRepository = (*userRepository_reflect_stub)(nil)

func (s userRepository_reflect_stub) Create(ctx context.Context, a0 User) (r0 User, err error) {
	err = s.caller("Create", ctx, []any{a0}, []any{&r0})
	return
}

func (s userRepository_reflect_stub) Delete(ctx context.Context, a0 uint64) (err error) {
	err = s.caller("Delete", ctx, []any{a0}, []any{})
	return
}

func (s userRepository_reflect_stub) FindAll(ctx context.Context) (r0 []User, err error) {
	err = s.caller("FindAll", ctx, []any{}, []any{&r0})
	return
}

func (s userRepository_reflect_stub) FindById(ctx context.Context, a0 uint64) (r0 User, err error) {
	err = s.caller("FindById", ctx, []any{a0}, []any{&r0})
	return
}

func (s userRepository_reflect_stub) Update(ctx context.Context, a0 User) (r0 User, err error) {
	err = s.caller("Update", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*User)(nil)

type __is_User[T ~struct {
	weaver.AutoMarshal
	Id        uint64 "json:\"id\""
	Name      string "json:\"name\""
	Email     string "json:\"email\""
	CreatedAt string "json:\"createdAt\""
	UpdatedAt string "json:\"updatedAt\""
}] struct{}

var _ __is_User[User]

func (x *User) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverMarshal: nil receiver"))
	}
	enc.Uint64(x.Id)
	enc.String(x.Name)
	enc.String(x.Email)
	enc.String(x.CreatedAt)
	enc.String(x.UpdatedAt)
}

func (x *User) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.Uint64()
	x.Name = dec.String()
	x.Email = dec.String()
	x.CreatedAt = dec.String()
	x.UpdatedAt = dec.String()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_User_42defd96(enc *codegen.Encoder, arg []User) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_User_42defd96(dec *codegen.Decoder) []User {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]User, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

// Size implementations.

// serviceweaver_size_User_0df39765 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_User_0df39765(x *User) int {
	size := 0
	size += 0
	size += 8
	size += (4 + len(x.Name))
	size += (4 + len(x.Email))
	size += (4 + len(x.CreatedAt))
	size += (4 + len(x.UpdatedAt))
	return size
}
