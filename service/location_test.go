package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"testing"
)

func TestGetLocationInfoByIPv4(t *testing.T) {
	pv4, err := GetLocationInfoByGeoDataTool("216.127.164.234")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", *pv4)
}

func TestGetLocationInfoByIpApi(t *testing.T) {
	api, err := GetLocationInfoByIpApi("208.95.112.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("info: %v \n", api)
	api2, err2 := GetLocationInfoByIpApi("2403:71c0:2000:a0c1:afc1::")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("info: %v \n", api2)

}

func TestGetIpgeolocationInfo(t *testing.T) {
	info, err := GetIpgeolocationInfo("208.95.112.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("info: %v \n", info)
}

func TestGetFlagEmoji(t *testing.T) {
	emoji := GetFlagEmoji("US")
	fmt.Println(emoji)
}

func TestGetFlagEmojiSimple(t *testing.T) {
	simple := GetFlagEmojiSimple("US")
	fmt.Println(simple)
}

func TestCheckNormalIpv6Address(t *testing.T) {
	address := CheckNormalIpv6Address("403:71c0:2000:a0c1:afc1::")
	fmt.Println(address)
}

func TestCheckNormalIpAddress(t *testing.T) {
	address := CheckNormalIpAddress("403:71c0:2000:a0c1:afc1::")
	fmt.Println(address)
	address = CheckNormalIpAddress("127.2.3.1")
	fmt.Println(address)
}

func TestCheckStrIsIpAddress(t *testing.T) {
	address := CheckStrIsIpAddress("403:71c0:2000:a0c1:afc1::")
	fmt.Println(address)
	address = CheckStrIsIpAddress("127.2.3.")
	fmt.Println(address)

}

func TestGoquery(t *testing.T) {
	var html = `<html xml:lang="en" lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
  <meta name="google-site-verification" content="3hmnzpoNToHe8lvKc9yFFWLnTwcKSf4UpDSqCGxGBPg"/>
  <meta name="DESCRIPTION" content="View my IP information: 216.127.164.234">
  <meta name="keywords" content="geo ip, geo map, geo ip tool, ip tool, ip2location, geo, region, city, map, data, ip, host, code, isp, ips, latitude, hostname, longitude, map, country, address, ip location">
  <link rel="alternate" href="https://geodatatool.com/en/" hreflang="en"/>
  <link rel="alternate" href="https://geodatatool.com/es/" hreflang="es"/>
  <link rel="alternate" href="https://geodatatool.com/ru/" hreflang="ru"/>
  <link rel="alternate" href="https://geodatatool.com/pt/" hreflang="pt"/>
  <link rel="alternate" href="https://geodatatool.com/de/" hreflang="de"/>
  <link rel="alternate" href="https://geodatatool.com/fr/" hreflang="fr"/>
  <link rel="alternate" href="https://geodatatool.com/zh/" hreflang="zh"/>
  <link rel="alternate" href="https://geodatatool.com/ja/" hreflang="ja"/>
  <link rel="alternate" href="https://geodatatool.com/it/" hreflang="it"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.css">
  <link rel="stylesheet" type="text/css" href="/static/css/geoip.css">
  <script type="text/javascript" src="https://maps.google.com/maps/api/js?sensor=true"></script>
  <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/gmaps.js"></script>
  <link rel="icon" type="image/png" href="/static/img/favicon.ico">
  <title>Geo Data Tool - View my IP information: 216.127.164.234 </title>
  <script data-ad-client="ca-pub-6560804155104041" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
</head>
<div class="top-header">
  <div class="container">
    <div class="row hidden-xs hidden-sm">
      <div class="col-md-4 col-xs-4 column">
        <div class="header-title" style="margin-top: 10px;">
          <a href="/">
            <img src="/static/img/logo.png">
          </a>
        </div>
      </div>
      <div class="col-md-8 col-xs-8">
        <nav class="navbar pull-right" role="navigation">
          <div class="container-fluid">
            <div class="collapse navbar-collapse desktop-collapse" id="bs-example-navbar-collapse-1">
              <ul class="nav navbar-nav">
                <li class="active">
                  <a href="/">View my IP information</a>
                </li>
                <li class="">
                  <a href="/en/ip_info">More info about IPs</a>
                </li>
                <li class="dropdown">
                  <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                    Language
                    <span class="caret"></span>
                  </a>
                  <ul class="dropdown-menu" role="menu">
                    <li>
                      <a href="/en" hreflang="en">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/us.gif">
                            </td>
                            <td>
                              <span>English</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/es" hreflang="es">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/es.gif">
                            </td>
                            <td>
                              <span>Spanish</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/ru" hreflang="ru">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/ru.gif">
                            </td>
                            <td>
                              <span>Russian</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/pt" hreflang="pt">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/pt.gif">
                            </td>
                            <td>
                              <span>Portuguese</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/fr" hreflang="fr">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/fr.gif">
                            </td>
                            <td>
                              <span>French</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/it" hreflang="it">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/it.gif">
                            </td>
                            <td>
                              <span>Italian</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/de" hreflang="de">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/de.gif">
                            </td>
                            <td>
                              <span>German</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/zh" hreflang="zh">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/cn.gif">
                            </td>
                            <td>
                              <span>Chinese</span>
                            </td>
                          </tr>
                          </tbody></table></a></li>
                    <li>
                      <a href="/ja" hreflang="ja">
                        <table>
                          <body>
                          <tr>
                            <td>
                              <img src="../static/img/flags/jp.gif">
                            </td>
                            <td>
                              <span>Japanese</span>
                            </td>
                          </tr>
                          </tbody></table></a></li></ul></li></ul></div></div></nav></div></div></div>
  <div class="mobile-header visible-xs visible-sm">
    <a href="/">
      <img src="/static/img/logo.png">
    </a>
  </div>
  <div class="mobile-navbar visible-xs visible-sm">
    <div class="mobile-navbar-container">
      <div class="container">
        <div class="row">
          <div class="col-sm-12">
            <div class="mobile-nav">
              <div class="row">
                <div class="col-sm-4 active">
                  <div class="mobile-item">
                    <a href="/">View my IP information</a>
                  </div>
                </div>
                <div class="col-sm-4 ">
                  <div class="mobile-item">
                    <a href="/en/ip_info">More info about IPs</a>
                  </div>
                </div>
                <div class="col-sm-4">
                  <div class="mobile-item mobile-item-dropdown">
                    <a href="#">Language</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="mobile-languages-container" style="display: none;">
      <div class="container">
        <div class="mobile-languages">
          <div class="row">
            <a href="/en">
              <div class="col-sm-6">
                <img src="../static/img/flags/us.gif">
                <span>English</span>
              </div>
            </a>
            <a href="/es">
              <div class="col-sm-6">
                <img src="../static/img/flags/es.gif">
                <span>Spanish</span>
              </div>
            </a>
          </div>
          <div class="row">
            <a href="/ru">
              <div class="col-sm-6">
                <img src="../static/img/flags/ru.gif">
                <span>Russian</span>
              </div>
            </a>
            <a href="/pt">
              <div class="col-sm-6">
                <img src="../static/img/flags/pt.gif">
                <span>Portuguese</span>
              </div>
            </a>
          </div>
          <div class="row">
            <a href="/fr">
              <div class="col-sm-6">
                <img src="../static/img/flags/fr.gif">
                <span>French</span>
              </div>
            </a>
            <a href="/it">
              <div class="col-sm-6">
                <img src="../static/img/flags/it.gif">
                <span>Italian</span>
              </div>
            </a>
          </div>
          <div class="row">
            <a href="/de">
              <div class="col-sm-6">
                <img src="../static/img/flags/de.gif">
                <span>German</span>
              </div>
            </a>
            <a href="/zh">
              <div class="col-sm-6">
                <img src="../static/img/flags/ch.gif">
                <span>Chinese</span>
              </div>
            </a>
          </div>
          <div class="row">
            <a href="/ja">
              <div class="col-sm-6">
                <img src="../static/img/flags/jp.gif">
                <span>Japanese</span>
              </div>
            </a>
            <div class="col-sm-6" style="display: none;"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div><script type="text/javascript">
  $(document).ready(function() {
    $('.mobile-item-dropdown').on('click', function() {
      var $this = $(this)
      $('.mobile-languages-container').slideToggle();
      $this.parents().find('.col-sm-4').last().toggleClass('active')
    });
  });
</script>
<script>
  (function(i, s, o, g, r, a, m) {
      i['GoogleAnalyticsObject'] = r;
      i[r] = i[r] || function() {
        (i[r].q = i[r].q || []).push(arguments)
      }
        ,
        i[r].l = 1 * new Date();
      a = s.createElement(o),
        m = s.getElementsByTagName(o)[0];
      a.async = 1;
      a.src = g;
      m.parentNode.insertBefore(a, m)
    }
  )(window, document, 'script', '//www.google-analytics.com/analytics.js', 'ga');

  ga('create', 'UA-82878-6', 'auto');
  ga('send', 'pageview');
</script>
<body>
<div class="container">
  <div class="main">
    <div class="row">
      <div class="col-md-12 column">
        <div class="geo-info-box">
          <div class="row">
            <div class="col-md-4 column">
              <div class="row hidden-xs hidden-sm">
                <div class="col-md-6">
                  <div class="info-box-banner">
                    <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
                    <!-- Geo Data Tool Top -->
                    <ins class="adsbygoogle" style="display:inline-block;width:240px;height:90px" data-ad-client="ca-pub-6560804155104041" data-ad-slot="2810736648"></ins>
                    <script>
                      (adsbygoogle = window.adsbygoogle || []).push({});
                    </script>
                  </div>
                </div>
              </div>
              <div class="footer-banner visible-xs visible-sm">
                <style>
                    .geo-ip-tool-responsive {
                        width: 300px;
                        height: 50px;
                    }

                    @media(min-width: 500px) {
                        .geo-ip-tool-responsive {
                            width: 300px;
                            height: 89px;
                        }
                    }

                    @media(min-width: 800px) {
                        .geo-ip-tool-responsive {
                            width: 300px;
                            height: 89px;
                        }
                    }
                </style>
                <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
                <!-- Geo IP Tool - Responsive -->
                <ins class="adsbygoogle geo-ip-tool-responsive" style="display:inline-block" data-ad-client="ca-pub-6560804155104041" data-ad-slot="6763896448"></ins>
                <script>
                  (adsbygoogle = window.adsbygoogle || []).push({});
                </script>
              </div>
              <div class="search-container hidden-xs hidden-sm">
                <form action="/en" method="GET" class="form-inline">
                  <div class="form-group">
                    <div class="row">
                      <div class="col-md-10">
                        <div class="input-group">
                          <div class="input-group-addon">Host/IP
                          </div>
                          <input type="text" name="ip" class="form-control">
                        </div>
                      </div>
                      <div class="col-md-1" style="width: 0; margin-right: -15px;">
                        <div class="input-group hidden-xs hidden-sm">
                          <button class="geo-btn geo-search-btn" type="submit" value="">
                            <span class="glyphicon glyphicon-search"></span>
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>
                </form>
              </div>
              <div class="mobile-search-container visible-xs visible-sm">
                <form action="/en" method="GET" class="form-inline">
                  <div class="row">
                    <div class="col-xs-3">
                      <div class="mobile-addon">Host/IP
                      </div>
                    </div>
                    <div class="col-xs-8">
                      <input type="text" name="ip" class="form-control">
                    </div>
                  </div>
                </form>
              </div>
              <div class="sidebar-data hidden-xs hidden-sm" style="margin-top: 5px;">
                <div class="data-item">
                  <span class="bold">Hostname:</span>
                  <span>
                                            <td>216.127.164.234</td>
                                        </span>
                </div>
                <div class="data-item">
                  <span class="bold">IP Address:</span>
                  <span>216.127.164.234</span>
                </div>
                <div class="data-item">
                  <span class="bold">Country:</span>
                  <span>
                                            <img src="../static/img/flags/us.gif">United States

                                        </span>
                </div>
                <div class="data-item">
                  <span class="bold">Country Code:</span>
                  <span>US ()</span>
                </div>
                <div class="data-item">
                  <span class="bold">Region:</span>
                  <span>California</span>
                </div>
                <div class="data-item">
                  <span class="bold">City:</span>
                  <span>Los Angeles</span>
                </div>
                <div class="data-item">
                  <span class="bold">Postal Code:</span>
                  <span>90001</span>
                </div>
                <div class="data-item">
                  <span class="bold">Latitude:</span>
                  <span>34.052230</span>
                </div>
                <div class="data-item">
                  <span class="bold">Longitude:</span>
                  <span>-118.243680</span>
                </div>
              </div>
              <div class="mobile-sidebar-data visible-xs visible-sm">
                <div class="data-item">
                  <span class="bold">Hostname:</span>
                  <span>
                                            <td>216.127.164.234</td>
                                        </span>
                </div>
                <div class="data-item">
                  <span class="bold">IP Address:</span>
                  <span>216.127.164.234</span>
                </div>
                <div class="data-item">
                  <span class="bold">Country:</span>
                  <span>
                                            <img src="../static/img/flags/us.gif">United States

                                        </span>
                </div>
                <div class="mobile-sidebar-hidden" style="display: none;">
                  <div class="data-item">
                    <span class="bold">Country Code:</span>
                    <span>US ()</span>
                  </div>
                  <div class="data-item">
                    <span class="bold">Region:</span>
                    <span>California</span>
                  </div>
                  <div class="data-item">
                    <span class="bold">City:</span>
                    <span>Los Angeles</span>
                  </div>
                  <div class="data-item">
                    <span class="bold">Latitude:</span>
                    <span>34.052230</span>
                  </div>
                  <div class="data-item">
                    <span class="bold">Longitude:</span>
                    <span>-118.243680</span>
                  </div>
                </div>
                <div class="mobile-button-display-hidden" data-value="0">
                  <span class="glyphicon glyphicon-plus"></span>
                </div>
              </div>
            </div>
            <div class="col-md-8 column">
              <div class="map-container">
                <div id="map" style="width: 100%;"></div>
              </div>
              <div class="footer-banner hidden-xs hidden-sm">
                <style>
                    .geo-ip-tool-responsive {
                        width: 320px;
                        height: 50px;
                    }

                    @media(min-width: 500px) {
                        .geo-ip-tool-responsive {
                            width: 468px;
                            height: 89px;
                        }
                    }

                    @media(min-width: 800px) {
                        .geo-ip-tool-responsive {
                            width: 615px;
                            height: 89px;
                        }
                    }
                </style>
                <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
                <!-- Geo IP Tool - Responsive -->
                <ins class="adsbygoogle geo-ip-tool-responsive" style="display:inline-block" data-ad-client="ca-pub-6560804155104041" data-ad-slot="6763896448"></ins>
                <script>
                  (adsbygoogle = window.adsbygoogle || []).push({});
                </script>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
  $(document).ready(function() {
    var map_height = $('.geo-info-box .col-md-4').height() - 83;
    $('#map').height(map_height);

    if ($(document).width() > 981) {
      var map_object = new GMaps({
        div: '#map',
        lat: 34.052230,
        lng: -118.243680,
        zoom: 4
      });

      var maker_country = "<div>Country: United States</div>";
      var marker_city = "<div>City: Los Angeles</div>";
      var marker_ip = "<div>IP Address: 216.127.164.234</div>";

      var marker_content = maker_country + marker_city + marker_ip;

      var marker = map_object.addMarker({
        lat: 34.052230,
        lng: -118.243680,
        title: "Los Angeles",
        infoWindow: {
          content: marker_content
        }
      });

      marker.infoWindow.open(map_object, marker);
    } else {
      var map_width = $(document).width() - 30
      url = GMaps.staticMapURL({
        size: [map_width, 300],
        lat: 34.052230,
        lng: -118.243680,
        markers: [{
          lat: 34.052230,
          lng: -118.243680
        }, ]
      });

      var img = $('<img/>').attr('src', url).appendTo('#map');
    }

    $('.mobile-button-display-hidden').on('click', function() {
      var $this = $(this);
      var value = parseInt($this.data('value'));

      $('.mobile-sidebar-hidden').slideToggle();

      if (value == 0) {
        $this.find('span').remove();
        $this.append('<span class="glyphicon glyphicon-minus"></span>')
        $this.data('value', 1)
      } else {
        $this.find('span').remove();
        $this.append('<span class="glyphicon glyphicon-plus"></span>')
        $this.data('value', 0)
      }
    });
  });
</script>
<div class="footer">
  <div class="container">
    <a href="https://twitter.com/share" class="twitter-share-button" data-url="https://www.geodatatool.com/" data-hashtags="GeoDataTool">Twittear
    </a>
    <script>
      !function(d, s, id) {
        var js, fjs = d.getElementsByTagName(s)[0];
        if (!d.getElementById(id)) {
          js = d.createElement(s);
          js.id = id;
          js.src = "//platform.twitter.com/widgets.js";
          fjs.parentNode.insertBefore(js, fjs);
        }
      }(document, "script", "twitter-wjs");
    </script>
    <g:plusone size="medium"></g:plusone>
    <script type="text/javascript">
      (function() {
          var po = document.createElement('script');
          po.type = 'text/javascript';
          po.async = true;
          po.src = 'https://apis.google.com/js/plusone.js';
          var s = document.getElementsByTagName('script')[0];
          s.parentNode.insertBefore(po, s);
        }
      )();
    </script>
    <iframe src="https://www.facebook.com/plugins/like.php?href=http%3A%2F%2Fwww.geoiptool.com%2F&amp;send=false&amp;layout=button_count&amp;width=150&amp;show_faces=false&amp;action=like&amp;colorscheme=light&amp;font&amp;height=21&amp;appId=223059641082996" scrolling="no" frameborder="0" style="border:none; overflow:hidden; width:150px; height:21px;" allowTransparency="true"></iframe>
    <a href="https://www.wiroos.com">
      <img src="/static/img/wiroos.png" class="pull-right" style="height: 25px;">
    </a>
  </div>
</div>
</body>
</html>
`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}
	var resultList []string

	dom.Find("body > div.container > div > div > div > div > div > div.col-md-4.column > div.sidebar-data.hidden-xs.hidden-sm > div > span:nth-child(2)").Each(func(i int, selection *goquery.Selection) {
		cleanStr := strings.TrimSpace(selection.Text())
		resultList = append(resultList, cleanStr)
		fmt.Println(cleanStr)
	})
	country := resultList[3]
	countryStrs := strings.Split(country, " ")
	country = countryStrs[0]
	resultList[3] = country
}
