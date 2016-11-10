// Package httpsnoop provides an easy way to capture http related metrics (i.e.
// response time, bytes written, and http status code) from your application.
//
// Doing this requires non-trivial wrapping of the http.ResponseWriter
// interface, which is also exposed for users interested in a more low-level
// API.
package httpsnoop
