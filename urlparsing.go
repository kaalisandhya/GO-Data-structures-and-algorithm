//8.url parsing
package main

	import "fmt"
import "net"
import "net/url"

	func main() {
// We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	    s := "postgres://user:pass@host.com:5432/path?k=v#f"
// Parse the URL and ensure there are no errors.
	    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
// Accessing the scheme is straightforward.
	    fmt.Println(u.Scheme)
// User contains all authentication info; call Username and Password on this for individual values.
	    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

// The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)
// Here we extract the path and the fragment after the #.
	    fmt.Println(u.Path)
    fmt.Println(u.Fragment)
// To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.
	    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v

Program exited.
package main
 
import (
    "fmt"
    "log"
    "net"
    "net/url"
    "strings"
)
 
func main() {
 
    var links = []string{
        "http://www.golangprograms.com/",
        "mailto:John.Mark@testing.com",
        "https://www.google.com/search?q=golang+print+string+10+times&oq=golang+print+string+10+times&aqs=chrome..69i57.8786j0j8&sourceid=chrome&ie=UTF-8",
        "urn:oasis:names:description:docbook:dtd:xml:4.1.2",
        "https://stackoverflow.com/jobs?med=site-ui&ref=jobs-tab",
        "ssh://mark@testing.com",
    }
    for _, link := range links {
 
        fmt.Println("URL:", link)
 
        u, err := url.Parse(link)
        if err != nil {
            log.Println(err)
            continue
        }
 
        parserURL(u)
        fmt.Println(strings.Repeat("#", 50))
        fmt.Println()
    }
}
 
func parserURL(u *url.URL) {
    fmt.Println("Scheme:", u.Scheme)
    if u.Opaque != "" {
        fmt.Println("Opaque:", u.Opaque)
    }
    if u.User != nil {
        fmt.Println("Username:", u.User.Username())
        if pwd, ok := u.User.Password(); ok {
            fmt.Println("Password:", pwd)
        }
    }
    if u.Host != "" {
        if host, port, err := net.SplitHostPort(u.Host); err == nil {
            fmt.Println("Host:", host)
            fmt.Println("Port:", port)
        } else {
            fmt.Println("Host:", u.Host)
        }
    }
    if u.Path != "" {
        fmt.Println("Path:", u.Path)
    }
    if u.RawQuery != "" {
        fmt.Println("RawQuery:", u.RawQuery)
        m, err := url.ParseQuery(u.RawQuery)
        if err == nil {
            for k, v := range m {
                fmt.Printf("Key: %q Values: %q\n", k, v)
            }
        }
    }
    if u.Fragment != "" {
        fmt.Println("Fragment:", u.Fragment)
    }
}
   