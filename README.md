# require
A simple require package to improve code stability and reduce bugs by strict requirement?

I recently read that bugs in more complex systems can be more easily avoided by explicitly validating inputs. 
I found this assumption interesting and decided to test it. 
To this end, I developed a library that provides the basic functionality for this purpose. 
I am sharing it here in case anyone else is interested.

***

To illustrate the concept more clearly, here is an example: 
*Let’s assume we are building a wrapper around a client. 
Our application will then use this wrapper. What we already know is that the client has a function `Send(ctx context.Context, req *client.Request)` and we know that our application will never use nil as request. So to ensure this, we might do something as follow.*
```go
func (c *ClientWrapper) Send(ctx context.Context, req *client.Request) (client.Response, error) {
  require.NotNil(req)
  return c.client.Send(ctx, req)
}
```
This would ensure, that we never call `Send` with nil request from somewhere in our application.
I know this example is not quite "interesting", but maybe you fine another meaningful use case for that.


**I first have to see for myself whether it really makes sense or is helpful.**
