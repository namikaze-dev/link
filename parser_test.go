package link_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/namikaze-dev/link"
)

func TestParseLinks(t *testing.T) {
	casese := []struct {
		name string
		in   string
		want []link.Link
	}{
		{
			name: "html containing a single anchor element",
			in:   html,
			want: []link.Link{
				{"/other-page", "A link to another page"},
			},
		},
		{
			name: "html containing a single anchor element with two children",
			in:   htmlA,
			want: []link.Link{
				{"/dog", "Something in a span Text not in a span Bold text!"},
			},
		},
		{
			name: "html containing two anchor elements",
			in:   htmlB,
			want: []link.Link{
				{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
				{"https://github.com/gophercises", "Gophercises is on Github!"},
			},
		},
		{
			name: "compound html containing three anchor elements",
			in:   htmlC,
			want: []link.Link{
				{"#", "Login"},
				{"/lost", "Lost? Need help?"},
				{"https://twitter.com/marcusolsson", "@marcusolsson"},
			},
		},
		{
			name: "html containing an anchor element with a comment",
			in:   htmlD,
			want: []link.Link{
				{"/dog-cat", "dog cat"},
			},
		},
	}

	for _, c := range casese {
		t.Run(c.name, func(t *testing.T) {
			got, err := link.Parse(strings.NewReader(c.in))
			if err != nil {
				t.Fatalf("unexpected error while Parsing Links: %v", err)
			}

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %#v want %#v", got, c.want)
			}
		})
	}
}

var (
	html = 
`
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

	htmlA = `
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
`
	htmlB = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`
	htmlC = `
<!DOCTYPE html>
<!--[if lt IE 7]> <html class="ie ie6 lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="ie ie7 lt-ie9 lt-ie8"        lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="ie ie8 lt-ie9"               lang="en"> <![endif]-->
<!--[if IE 9]>    <html class="ie ie9"                      lang="en"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en" class="no-ie">
<!--<![endif]-->

<head>
  <title>Gophercises - Coding exercises for budding gophers</title>
</head>

<body>
  <section class="header-section">
    <div class="jumbo-content">
      <div class="pull-right login-section">
        Already have an account?
        <a href="#" class="btn btn-login">Login <i class="fa fa-sign-in" aria-hidden="true"></i></a>
      </div>
      <center>
        <img src="https://gophercises.com/img/gophercises_logo.png" style="max-width: 85%; z-index: 3;">
        <h1>coding exercises for budding gophers</h1>
        <br/>
        <form action="/do-stuff" method="post">
          <div class="input-group">
            <input type="email" id="drip-email" name="fields[email]" class="btn-input" placeholder="Email Address" required>
            <button class="btn btn-success btn-lg" type="submit">Sign me up!</button>
            <a href="/lost">Lost? Need help?</a>
          </div>
        </form>
        <p class="disclaimer disclaimer-box">Gophercises is 100% FREE, but is currently in beta. There will be bugs, and things will be changing significantly over the coming weeks.</p>
      </center>
    </div>
  </section>
  <section class="footer-section">
    <div class="row">
      <div class="col-md-6 col-md-offset-1 vcenter">
        <div class="quote">
          "Success is no accident. It is hard work, perseverance, learning, studying, sacrifice and most of all, love of what you are doing or learning to do." - Pele
        </div>
      </div>
      <div class="col-md-4 col-md-offset-0 vcenter">
        <center>
          <img src="https://gophercises.com/img/gophercises_lifting.gif" style="width: 80%">
          <br/>
          <br/>
        </center>
      </div>
    </div>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <center>
          <p class="disclaimer">
            Artwork created by Marcus Olsson (<a href="https://twitter.com/marcusolsson">@marcusolsson</a>), animated by Jon Calhoun (that's me!), and inspired by the original Go Gopher created by Renee French.
          </p>
        </center>
      </div>
    </div>
  </section>
</body>
</html>
`
	htmlD = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`
)
