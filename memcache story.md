Consegui codigo del siguiente repo:

- [link al repo](https://github.com/GoogleCloudPlatform/golang-samples/blob/22d8055a29603cf7f9e459522f7dbb5d70d726f8/docs/appengine/memcache/memcache.go)

El primer problema es que al pegarle al endpoint asociado con ese codigo, recibo el error: `error adding item: not an App Engine context`. Según leí, no es buena idea mezclar paquetes `"cloud.google.com/go/` con `"google.golang.org/appengine/`, hay que usar uno u otro ([fuente aqui](https://github.com/googleapis/google-cloud-go/issues/460)).

La solucion fue, finalmente, agregar una invocación a `appengine.Main()` que supuestamente manejara el tema de contexto. Pero hecho eso, tampoco anduvo la cosa. La app arrancaba, pero enseguida incurría en el siguiente problema:
```
Defaulting to port 8080
http.ListenAndServe: listen tcp :8080: bind: address already in use
```
Al investigar, se ve que `appengine.Main()` invoca un `http.ListenAndServe`, por lo que es lógico que la cosa colapse si hay _dos_ `ListenAndServe` queriendo ejecutarse, más si es _en el mismo puerto_. Dejo el código de dicho `Main()`:
```go
func Main() {
	MainPath = filepath.Dir(findMainPath())
	installHealthChecker(http.DefaultServeMux)

	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	host := ""
	if IsDevAppServer() {
		host = "127.0.0.1"
	}
	if err := http.ListenAndServe(host+":"+port, http.HandlerFunc(handleHTTP)); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
```
Comentando mi `ListenAndServe()` para que no haya dos invocaciones a lo mismo con puertos distintos, llego a este error:
```go
panic: Metadata fetch failed for 'instance/attributes/gae_project': Get "http://metadata/computeMetadata/v1/instance/attributes/gae_project": dial tcp: lookup metadata: Temporary failure in name resolution
```
y finalmente un exit status 2 descrito de esta forma:
```
goroutine 42 [running]:
google.golang.org/appengine/internal.mustGetMetadata(0xc1bd99, 0x1f, 0x0, 0x0, 0xc0000a9100)
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/metadata.go:34 +0x148
google.golang.org/appengine/internal.partitionlessAppID(0xc14103, 0x15)
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/identity_vm.go:112 +0xa5
google.golang.org/appengine/internal.DefaultTicket.func1()
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/api.go:297 +0x9c
sync.(*Once).doSlow(0x11624b8, 0xc3aac0)
	/usr/local/go/src/sync/once.go:66 +0xec
sync.(*Once).Do(...)
	/usr/local/go/src/sync/once.go:57
google.golang.org/appengine/internal.DefaultTicket(0xc0000e30e0, 0xabdc60)
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/api.go:292 +0x65
google.golang.org/appengine/internal.Call(0xce6220, 0xc0000e30e0, 0xc0d10f, 0xa, 0xc082fb, 0x5, 0xce3360, 0xc0000a91c0, 0xce2560, 0xc000031e60, ...)
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/api.go:499 +0x96a
google.golang.org/appengine/internal.(*context).flushLog(0xc0000a4480, 0xc3c601, 0x0)
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/api.go:644 +0x48c
google.golang.org/appengine/internal.handleHTTP.func1(0xc0000c0180, 0xc0000a4480)
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/api.go:141 +0x58
created by google.golang.org/appengine/internal.handleHTTP
	/home/rozanecm/go/pkg/mod/google.golang.org/appengine@v1.6.7/internal/api.go:137 +0x397
exit status 2
```
