package server

/* MetaPack contains metadata and the client identifier of
the client that send this metadata. */
type MetaPack struct {
    // The metadata in UTF-8 format
    Data string
    // The client identifier generated by the http server
    ID *ClientID
    // Indicates if we've seen this package before
    Seen bool
}

/* DataPack contains a data slice and a pointer to the client that
send this data to us.

This is generated by a reader and finally processed by the manager main
loop */
type DataPack struct {
    // The audio data
    Data []byte
    // A pointer to the client that send the data
    Client *Client
}

/* ErrPack contains an error and a pointer to the client that generated
this error. */
type ErrPack struct {
    // The error that occurred
    Err error
    // A pointer to the client that send the data
    Client *Client
}