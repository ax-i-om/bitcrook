package load

import (
	"github.com/audioo/goseek/helpers/ent"
)

// NoRedirSites ...
func NoRedirSites(userres string) []ent.Website {
	var arr []ent.Website

	arr = append(arr, ent.Website{Title: "FACEBOOK", Domain: "https://www.facebook.com/" + userres})
	arr = append(arr, ent.Website{Title: "BLOGSPOT", Domain: "https://" + userres + ".blogspot.com"})
	arr = append(arr, ent.Website{Title: "WORDPRESS", Domain: "https://" + userres + ".wordpress.com"})

	return arr
}

// RedirSites ...
func RedirSites(userres string) []ent.Website {
	var arr []ent.Website

	arr = append(arr, ent.Website{Title: "PINTEREST", Domain: "https://www.pinterest.com/" + userres})
	arr = append(arr, ent.Website{Title: "GITHUB", Domain: "https://www.github.com/" + userres})
	arr = append(arr, ent.Website{Title: "TUMBLR", Domain: "https://" + userres + ".tumblr.com"})
	arr = append(arr, ent.Website{Title: "FLICKR", Domain: "https://www.flickr.com/people/" + userres})
	arr = append(arr, ent.Website{Title: "VIMEO", Domain: "https://vimeo.com/" + userres})
	arr = append(arr, ent.Website{Title: "SOUNDCLOUD", Domain: "https://soundcloud.com/" + userres})
	arr = append(arr, ent.Website{Title: "DISQUS", Domain: "https://disqus.com/by/" + userres})
	arr = append(arr, ent.Website{Title: "MEDIUM", Domain: "https://medium.com/@" + userres})
	arr = append(arr, ent.Website{Title: "DEVIANTART", Domain: "https://" + userres + ".deviantart.com"})
	arr = append(arr, ent.Website{Title: "VK", Domain: "https://vk.com/" + userres})
	arr = append(arr, ent.Website{Title: "ABOUTME", Domain: "https://about.me/" + userres})
	arr = append(arr, ent.Website{Title: "FLIPBOARD", Domain: "https://flipboard.com/@" + userres})
	arr = append(arr, ent.Website{Title: "SLIDESHARE", Domain: "https://slideshare.net/" + userres})
	arr = append(arr, ent.Website{Title: "SPOTIFY", Domain: "https://open.spotify.com/user/" + userres})
	arr = append(arr, ent.Website{Title: "SCRIBD", Domain: "https://www.scribd.com/" + userres})
	arr = append(arr, ent.Website{Title: "PATREON", Domain: "https://www.patreon.com/" + userres})
	arr = append(arr, ent.Website{Title: "BITBUCKET", Domain: "https://bitbucket.org/" + userres})
	arr = append(arr, ent.Website{Title: "DAILYMOTION", Domain: "https://www.dailymotion.com/" + userres})
	arr = append(arr, ent.Website{Title: "ETSY", Domain: "https://www.etsy.com/shop/" + userres})
	arr = append(arr, ent.Website{Title: "CASHAPP", Domain: "https://cash.app/$" + userres})
	arr = append(arr, ent.Website{Title: "BEHANCE", Domain: "https://www.behance.net/" + userres})
	arr = append(arr, ent.Website{Title: "GOODREADS", Domain: "https://www.goodreads.com/" + userres})
	arr = append(arr, ent.Website{Title: "INSTRUCTABLES", Domain: "https://www.instructables.com/member/" + userres})
	arr = append(arr, ent.Website{Title: "KEYBASE", Domain: "https://keybase.io/" + userres})
	arr = append(arr, ent.Website{Title: "KONGREGATE", Domain: "https://kongregate.com/accounts/" + userres})
	arr = append(arr, ent.Website{Title: "LIVEJOURNAL", Domain: "https://" + userres + ".livejournal.com"})
	arr = append(arr, ent.Website{Title: "LASTFM", Domain: "https://last.fm/user/" + userres})
	arr = append(arr, ent.Website{Title: "DRIBBBLE", Domain: "https://dribbble.com/" + userres})
	arr = append(arr, ent.Website{Title: "CODECADEMY", Domain: "https://www.codecademy.com/profiles/" + userres})
	arr = append(arr, ent.Website{Title: "GRAVATAR", Domain: "https://en.gravatar.com/" + userres})
	arr = append(arr, ent.Website{Title: "PASTEBIN", Domain: "https://pastebin.com/u/" + userres})
	arr = append(arr, ent.Website{Title: "ROBLOX", Domain: "https://www.roblox.com/user.aspx?username=" + userres})
	arr = append(arr, ent.Website{Title: "GUMROAD", Domain: "https://www.gumroad.com/" + userres})
	arr = append(arr, ent.Website{Title: "NEWGROUNDS", Domain: "https://" + userres + ".newgrounds.com"})
	arr = append(arr, ent.Website{Title: "WATTPAD", Domain: "https://www.wattpad.com/user/" + userres})
	arr = append(arr, ent.Website{Title: "TRAKT", Domain: "https://www.trakt.tv/users/" + userres})
	arr = append(arr, ent.Website{Title: "BUZZFEED", Domain: "https://buzzfeed.com/" + userres})
	arr = append(arr, ent.Website{Title: "TRIPADVISOR", Domain: "https://www.tripadvisor.com/Profile/" + userres})
	arr = append(arr, ent.Website{Title: "HUBPAGES", Domain: "https://" + userres + ".hubpages.com"})
	arr = append(arr, ent.Website{Title: "BLIPFM", Domain: "https://blip.fm/" + userres})
	arr = append(arr, ent.Website{Title: "WIKIPEDIA", Domain: "https://en.wikipedia.org/wiki/User:" + userres})
	arr = append(arr, ent.Website{Title: "CODEMENTOR", Domain: "https://www.codementor.io/" + userres})
	arr = append(arr, ent.Website{Title: "REVERBNATION", Domain: "https://www.reverbnation.com/" + userres})
	arr = append(arr, ent.Website{Title: "DESIGNSPIRATION", Domain: "https://www.designspiration.com/" + userres})
	arr = append(arr, ent.Website{Title: "BANDCAMP", Domain: "https://www.bandcamp.com/" + userres})
	arr = append(arr, ent.Website{Title: "COLOURLOVERS", Domain: "https://www.colourlovers.com/lover/" + userres})
	arr = append(arr, ent.Website{Title: "IFTTT", Domain: "https://www.ifttt.com/p/" + userres})
	arr = append(arr, ent.Website{Title: "SLACK", Domain: "https://" + userres + ".slack.com"})
	arr = append(arr, ent.Website{Title: "ELLO", Domain: "https://ello.co/" + userres})

	return arr
}
