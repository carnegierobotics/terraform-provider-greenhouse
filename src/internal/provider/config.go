package greenhouse

import (
  "context"
  "fmt"
  "net/http"
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "golang.org/x/oauth2"
)

type Config struct {
  Url   string
  Token string
}

func NewConfig(d *schema.ResourceData) (*Config, error) {
  token, ok := d.GetOk("token")
  if !ok {
    return nil, fmt.Errorf("You must supply a token.")
  }
  url, ok := d.GetOk("url")
  if !ok {
    return nil, fmt.Errorf("You must supply a URL.")
  }
  c := &Config {
    Token: token.(string),
    Url:   url.(string),
  }
  return c, nil
}

func (c *Config) AuthenticatedHTTPClient() *http.Client {

  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: c.Token},
  )
  client := oauth2.NewClient(ctx, ts)

  return client
}

func (c *Config) Client() interface{} {
  client := c.AuthenticatedHTTPClient()
  return client
}
