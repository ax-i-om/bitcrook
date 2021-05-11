package load

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/maraudery/goseek/internal/cli"
	"github.com/maraudery/goseek/internal/http"
	"github.com/maraudery/goseek/pkg/ent"
)

// NoRedirSites ...
func NoRedirSites(userres string) []ent.Website {
	var arr []ent.Website

	arr = append(arr, ent.Website{Title: "FACEBOOK", Domain: "https://www.facebook.com/" + userres, Delete: "https://www.facebook.com/help/delete_account?rdrhc"})
	arr = append(arr, ent.Website{Title: "BLOGSPOT", Domain: "https://" + userres + ".blogspot.com", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "WORDPRESS", Domain: "https://" + userres + ".wordpress.com", Delete: "N/A"})

	return arr
}

// RedirSites ...
func RedirSites(userres string) []ent.Website {
	var arr []ent.Website

	arr = append(arr, ent.Website{Title: "PINTEREST", Domain: "https://www.pinterest.com/" + userres, Delete: "https://pinterest.com/settings/"})
	arr = append(arr, ent.Website{Title: "GITHUB", Domain: "https://www.github.com/" + userres, Delete: "https://github.com/settings/admin", Extra: githubExtra(userres)})
	arr = append(arr, ent.Website{Title: "TUMBLR", Domain: "https://" + userres + ".tumblr.com", Delete: "https://www.tumblr.com/account/delete"})
	arr = append(arr, ent.Website{Title: "FLICKR", Domain: "https://www.flickr.com/people/" + userres, Delete: "http://www.flickr.com/profile_delete.gne"})
	arr = append(arr, ent.Website{Title: "VIMEO", Domain: "https://vimeo.com/" + userres, Delete: "https://vimeo.com/settings/goodbye/forever"})
	arr = append(arr, ent.Website{Title: "SOUNDCLOUD", Domain: "https://soundcloud.com/" + userres, Delete: "http://soundcloud.com/settings/account#delete-user"})
	arr = append(arr, ent.Website{Title: "DISQUS", Domain: "https://disqus.com/by/" + userres, Delete: "http://disqus.com/pages/dashboard/#account", Extra: disqusExtra(userres)})
	arr = append(arr, ent.Website{Title: "MEDIUM", Domain: "https://medium.com/@" + userres, Delete: "https://medium.com/me/settings"})
	arr = append(arr, ent.Website{Title: "DEVIANTART", Domain: "https://" + userres + ".deviantart.com", Delete: "https://www.deviantart.com/settings/deactivation"})
	arr = append(arr, ent.Website{Title: "VK", Domain: "https://vk.com/" + userres, Delete: "http://vk.com/settings?act=deactivate"})
	arr = append(arr, ent.Website{Title: "ABOUTME", Domain: "https://about.me/" + userres, Delete: "https://about.me/account"})
	arr = append(arr, ent.Website{Title: "FLIPBOARD", Domain: "https://flipboard.com/@" + userres, Delete: "https://flipboard.com/support/"})
	arr = append(arr, ent.Website{Title: "SLIDESHARE", Domain: "https://slideshare.net/" + userres, Delete: "N/A"})
	arr = append(arr, ent.Website{Title: "SCRIBD", Domain: "https://www.scribd.com/" + userres, Delete: "http://www.scribd.com/account_settings/preferences"})
	arr = append(arr, ent.Website{Title: "PATREON", Domain: "https://www.patreon.com/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "BITBUCKET", Domain: "https://bitbucket.org/" + userres, Delete: "https://bitbucket.org/account/"})
	arr = append(arr, ent.Website{Title: "DAILYMOTION", Domain: "https://www.dailymotion.com/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "ETSY", Domain: "https://www.etsy.com/shop/" + userres, Delete: "http://www.etsy.com/help/article/53"})
	arr = append(arr, ent.Website{Title: "CASHAPP", Domain: "https://cash.app/$" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "BEHANCE", Domain: "https://www.behance.net/" + userres, Delete: "http://www.behance.net/account/privacy"})
	arr = append(arr, ent.Website{Title: "GOODREADS", Domain: "https://www.goodreads.com/" + userres, Delete: "http://www.goodreads.com/user/destroy"})
	arr = append(arr, ent.Website{Title: "INSTRUCTABLES", Domain: "https://www.instructables.com/member/" + userres, Delete: "Email service@instructables.com"})
	arr = append(arr, ent.Website{Title: "KEYBASE", Domain: "https://keybase.io/" + userres, Delete: "TODO", Extra: keybaseExtra(userres)})
	arr = append(arr, ent.Website{Title: "KONGREGATE", Domain: "https://kongregate.com/accounts/" + userres, Delete: "http://www.kongregate.com/forums/7/topics/241772?page=1#posts-5207538"})
	arr = append(arr, ent.Website{Title: "LIVEJOURNAL", Domain: "https://" + userres + ".livejournal.com", Delete: "http://www.livejournal.com/accountstatus.bml"})
	arr = append(arr, ent.Website{Title: "LASTFM", Domain: "https://last.fm/user/" + userres, Delete: "http://www.last.fm/settings/account"})
	arr = append(arr, ent.Website{Title: "DRIBBBLE", Domain: "https://dribbble.com/" + userres, Delete: "http://dribbble.com/account"})
	arr = append(arr, ent.Website{Title: "CODECADEMY", Domain: "https://www.codecademy.com/profiles/" + userres, Delete: "http://help.codecademy.com/customer/portal/articles/1394081-how-do-i-delete-my-account-"})
	arr = append(arr, ent.Website{Title: "GRAVATAR", Domain: "https://en.gravatar.com/" + userres, Delete: "N/A", Extra: gravatarExtra(userres)})
	arr = append(arr, ent.Website{Title: "PASTEBIN", Domain: "https://pastebin.com/u/" + userres, Delete: "N/A"})
	arr = append(arr, ent.Website{Title: "ROBLOX", Domain: "https://www.roblox.com/user.aspx?username=" + userres, Delete: "N/A", Extra: robloxExtra(userres)})
	arr = append(arr, ent.Website{Title: "GUMROAD", Domain: "https://www.gumroad.com/" + userres, Delete: "https://gumroad.com/settings"})
	arr = append(arr, ent.Website{Title: "NEWGROUNDS", Domain: "https://" + userres + ".newgrounds.com", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "WATTPAD", Domain: "https://www.wattpad.com/user/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "TRAKT", Domain: "https://www.trakt.tv/users/" + userres, Delete: "http://trakt.tv/settings/data"})
	arr = append(arr, ent.Website{Title: "BUZZFEED", Domain: "https://buzzfeed.com/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "TRIPADVISOR", Domain: "https://www.tripadvisor.com/Profile/" + userres, Delete: "https://www.tripadvisorsupport.com/hc/en-us/articles/200615117"})
	arr = append(arr, ent.Website{Title: "HUBPAGES", Domain: "https://" + userres + ".hubpages.com", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "BLIPFM", Domain: "https://blip.fm/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "WIKIPEDIA", Domain: "https://en.wikipedia.org/wiki/User:" + userres, Delete: "https://en.wikipedia.org/wiki/Wikipedia:FAQ#How_do_I_change_my_username.2Fdelete_my_account.3F"})
	arr = append(arr, ent.Website{Title: "CODEMENTOR", Domain: "https://www.codementor.io/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "REVERBNATION", Domain: "https://www.reverbnation.com/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "DESIGNSPIRATION", Domain: "https://www.designspiration.com/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "BANDCAMP", Domain: "https://www.bandcamp.com/" + userres, Delete: "https://bandcamp.com/settings?pane=fan"})
	arr = append(arr, ent.Website{Title: "COLOURLOVERS", Domain: "https://www.colourlovers.com/lover/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "IFTTT", Domain: "https://www.ifttt.com/p/" + userres, Delete: "https://ifttt.com/settings/deactivate"})
	arr = append(arr, ent.Website{Title: "SLACK", Domain: "https://" + userres + ".slack.com", Delete: "https://my.slack.com/account/settings"})
	arr = append(arr, ent.Website{Title: "ELLO", Domain: "https://ello.co/" + userres, Delete: "https://ello.co/delete-account"})
	arr = append(arr, ent.Website{Title: "7CUPS", Domain: "https://www.7cups.com/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "INDEPENDENT ACADEMIA", Domain: "https://independent.academia.edu/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "ALIK", Domain: "https://www.alik.cz/u/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "AUDIO JUNGLE", Domain: "https://audiojungle.net/user/" + userres + "", Delete: "https://help.market.envato.com/hc/en-us/articles/202500394-How-Do-I-Close-My-Account-"})
	arr = append(arr, ent.Website{Title: "BOOK CROSSING", Domain: "https://www.bookcrossing.com/mybookshelf/" + userres + "/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "BUY ME A COFFEE", Domain: "https://buymeacoff.ee/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "CNET", Domain: "https://www.cnet.com/profiles/" + userres + "/", Delete: "https://cbsi.secure.force.com/CBSi/submitcase?template=template_cnet&referer=cnet.com&cfs=SFS_1"})
	arr = append(arr, ent.Website{Title: "CARBONMADE", Domain: "https://" + userres + ".carbonmade.com", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Chess", Domain: "https://www.chess.com/member/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "CodeWars", Domain: "https://www.codewars.com/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Countable", Domain: "https://www.countable.us/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "DevTo", Domain: "https://dev.to/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Discogs", Domain: "https://www.discogs.com/user/" + userres + "", Delete: "http://www.discogs.com/users/delete"})
	arr = append(arr, ent.Website{Title: "Eyeem", Domain: "https://www.eyeem.com/u/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "F3", Domain: "https://f3.cool/" + userres + "/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Fortnite Tracker", Domain: "https://fortnitetracker.com/profile/all/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "FreeSound", Domain: "https://freesound.org/people/" + userres + "/", Delete: "http://www.freesound.org/home/delete/"})
	arr = append(arr, ent.Website{Title: "GameSpot", Domain: "https://www.gamespot.com/profile/" + userres + "/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Giphy", Domain: "https://giphy.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Gitlab", Domain: "https://gitlab.com/" + userres + "", Delete: "TODO", Extra: gitlabExtra(userres)})
	arr = append(arr, ent.Website{Title: "GuruShots", Domain: "https://gurushots.com/" + userres + "/photos", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Hackaday", Domain: "https://hackaday.io/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "HackerOne", Domain: "https://hackerone.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Houzz", Domain: "https://houzz.com/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Issuu", Domain: "https://issuu.com/" + userres + "", Delete: "https://issuu.com/home/settings/billing", Extra: issuuExtra(userres)})
	arr = append(arr, ent.Website{Title: "Itch", Domain: "https://" + userres + ".itch.io/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Jimdosite", Domain: "https://" + userres + ".jimdosite.com", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "LeetCode", Domain: "https://leetcode.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Letterboxd", Domain: "https://letterboxd.com/" + userres + "", Delete: "http://letterboxd.com/user/disableaccount/"})
	arr = append(arr, ent.Website{Title: "Lichess", Domain: "https://lichess.org/@/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Memrise", Domain: "https://www.memrise.com/user/" + userres + "/", Delete: "http://www.memrise.com/settings/deactivate/"})
	arr = append(arr, ent.Website{Title: "Munzee", Domain: "https://www.munzee.com/m/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "My Anime List", Domain: "https://myanimelist.net/profile/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "My Mini Factory", Domain: "https://www.myminifactory.com/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "MySpace", Domain: "https://myspace.com/" + userres + "", Delete: "https://myspace.com/settings/profile"})
	arr = append(arr, ent.Website{Title: "Periscope", Domain: "https://www.periscope.tv/" + userres + "/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Pink Bike", Domain: "https://www.pinkbike.com/u/" + userres + "/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Pokemon Showdown", Domain: "https://pokemonshowdown.com/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Product Hunt", Domain: "https://www.producthunt.com/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "PromoDJ", Domain: "http://promodj.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "PyPi", Domain: "https://pypi.org/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Rajce", Domain: "https://" + userres + ".rajce.idnes.cz/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Rate Your Music", Domain: "https://rateyourmusic.com/~" + userres + "", Delete: "http://rateyourmusic.com/account/delete"})
	arr = append(arr, ent.Website{Title: "Red Bubble", Domain: "https://www.redbubble.com/people/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Repl.it", Domain: "https://repl.it/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Sbazar", Domain: "https://www.sbazar.cz/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Shitpost Bot", Domain: "https://www.shitpostbot.com/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Smule", Domain: "https://www.smule.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "SourceForge", Domain: "https://sourceforge.net/u/" + userres + "", Delete: "https://sourceforge.net/auth/disable/"})
	arr = append(arr, ent.Website{Title: "Speedrun", Domain: "https://speedrun.com/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Roberts Space Industries", Domain: "https://robertsspaceindustries.com/citizens/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Tellonym", Domain: "https://tellonym.me/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "TikTok", Domain: "https://tiktok.com/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Trading View", Domain: "https://www.tradingview.com/u/" + userres + "/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Ultimate Guitar", Domain: "https://ultimate-guitar.com/u/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Unsplash", Domain: "https://unsplash.com/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "VSCO", Domain: "https://vsco.co/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Vero", Domain: "https://vero.co/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "WarriorForum", Domain: "https://www.warriorforum.com/members/" + userres + ".html", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "We Heart It", Domain: "https://weheartit.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Windy", Domain: "https://community.windy.com/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "XBox", Domain: "https://xboxgamertag.com/search/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Zhihu", Domain: "https://www.zhihu.com/people/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "All My Links", Domain: "https://allmylinks.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Couch Surfing", Domain: "https://www.couchsurfing.com/people/" + userres + "", Delete: "https://www.couchsurfing.org/delete_profile.html?delete=1"})
	arr = append(arr, ent.Website{Title: "DailyKOs", Domain: "https://www.dailykos.com/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Dating", Domain: "http://dating.ru/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Drive2", Domain: "https://www.drive2.ru/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "FL", Domain: "https://www.fl.ru/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Geocaching", Domain: "https://www.geocaching.com/p/default.aspx?u=" + userres + "", Delete: "http://support.groundspeak.com/index.php?pg=kb.page&id=102"})
	arr = append(arr, ent.Website{Title: "Gfycat", Domain: "https://gfycat.com/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Hackster", Domain: "https://www.hackster.io/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Irecommend", Domain: "https://irecommend.ru/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "JBZD", Domain: "https://jbzd.com.pl/uzytkownik/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "JeuxVideo", Domain: "http://www.jeuxvideo.com/profil/" + userres + "?mode=infos", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Kwork", Domain: "https://kwork.ru/user/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "Livelib", Domain: "https://www.livelib.ru/reader/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "MOIKRUG", Domain: "https://moikrug.ru/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "NAIRALAND", Domain: "https://www.nairaland.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "NN", Domain: "https://" + userres + ".www.nn.ru/", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "NOTE", Domain: "https://note.com/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "NPMJS", Domain: "https://www.npmjs.com/~" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "OSU", Domain: "https://osu.ppy.sh/users/" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "PIKABU", Domain: "https://pikabu.ru/@" + userres + "", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "TOSTER", Domain: "https://www.toster.ru/user/" + userres + "/answers", Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "PORNHUB", Domain: "https://pornhub.com/users/" + userres})
	arr = append(arr, ent.Website{Title: "XVIDEOS", Domain: "https://xvideos.com/profiles/" + userres})
	arr = append(arr, ent.Website{Title: "YOUPORN", Domain: "https://youporn.com/uservids/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "XHAMSTER", Domain: "https://xhamster.com/users/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "BONGACAMS", Domain: "https://pt.bongacams.com/profile/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "ROYALCAMS", Domain: "https://royalcams.com/profile/" + userres, Delete: "TODO"})
	arr = append(arr, ent.Website{Title: "CHATURBATE", Domain: "https://chaturbate.com/" + userres, Delete: "TODO"})

	return arr
}

/* Begin Extras */

type github struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func githubExtra(userres string) string {
	var resp = http.GetReq("https://api.github.com/users/" + userres)
	var s = new(github)
	err := json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", strconv.Itoa(s.ID), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Node ID: ", s.NodeID, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Avatar URL: ", s.AvatarURL, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Gravatar ID: ", s.GravatarID, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Type: ", s.Type, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Site Admin: ", strconv.FormatBool(s.SiteAdmin), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Name: ", s.Name, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Company: ", s.Company, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Blog: ", s.Blog, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Location: ", s.Location, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Email: ", s.Email, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Hireable: ", strconv.FormatBool(s.Hireable), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Bio: ", s.Bio, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Twitter Username: ", s.TwitterUsername, firstSuccess)
	return r
}

type robloxone struct {
	ID          int         `json:"Id"`
	Username    string      `json:"Username"`
	Avataruri   interface{} `json:"AvatarUri"`
	Avatarfinal bool        `json:"AvatarFinal"`
	Isonline    bool        `json:"IsOnline"`
}

type robloxtwo struct {
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Isbanned    bool      `json:"isBanned"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Displayname string    `json:"displayName"`
}

func robloxExtra(userres string) string {
	var resp = http.GetReq("https://api.roblox.com/users/get-by-username?username=" + userres)
	var spre = new(robloxone)
	err := json.Unmarshal([]byte(resp), &spre)
	if err != nil {
		return ""
	}
	var resp2 = http.GetReq("https://users.roblox.com/v1/users/" + strconv.Itoa(spre.ID))
	var s = new(robloxtwo)
	err = json.Unmarshal([]byte(resp2), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", strconv.Itoa(s.ID), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Creation Date: ", s.Created.String(), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Is Online: ", strconv.FormatBool(spre.Isonline), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Is Banned: ", strconv.FormatBool(s.Isbanned), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Description: ", s.Description, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Name: ", s.Name, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Display Name: ", s.Displayname, firstSuccess)
	return r
}

type gravatar struct {
	Entry []struct {
		ID                string `json:"id"`
		Hash              string `json:"hash"`
		Requesthash       string `json:"requestHash"`
		Profileurl        string `json:"profileUrl"`
		Preferredusername string `json:"preferredUsername"`
		Thumbnailurl      string `json:"thumbnailUrl"`
		Photos            []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"photos"`
		Displayname string        `json:"displayName"`
		Urls        []interface{} `json:"urls"`
	} `json:"entry"`
}

func gravatarExtra(userres string) string {
	var resp = http.GetReq("https://en.gravatar.com/" + userres + ".json")
	var s = new(gravatar)
	err := json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", s.Entry[0].ID, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Hash: ", s.Entry[0].Hash, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Request Hash: ", s.Entry[0].Requesthash, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Preferred Username: ", s.Entry[0].Preferredusername, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Thumbnail URL: ", s.Entry[0].Thumbnailurl, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Display Name: ", s.Entry[0].Displayname, firstSuccess)
	return r
}

type keybase struct {
	Status struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"status"`
	Them []struct {
		ID     string `json:"id"`
		Basics struct {
			Username      string `json:"username"`
			Ctime         int    `json:"ctime"`
			Mtime         int    `json:"mtime"`
			IDVersion     int    `json:"id_version"`
			TrackVersion  int    `json:"track_version"`
			LastIDChange  int    `json:"last_id_change"`
			UsernameCased string `json:"username_cased"`
			Status        int    `json:"status"`
			Salt          string `json:"salt"`
			EldestSeqno   int    `json:"eldest_seqno"`
		} `json:"basics"`
		PublicKeys struct {
			AllBundles []string `json:"all_bundles"`
			Subkeys    []string `json:"subkeys"`
			Sibkeys    []string `json:"sibkeys"`
			Families   struct {
				Zero120B870B40Ba525Cc168433Bcaadcaf41904Af043455779D619C73Cd8B5A72C1De90A []string `json:"0120b870b40ba525cc168433bcaadcaf41904af043455779d619c73cd8b5a72c1de90a"`
			} `json:"families"`
			EldestKid            string        `json:"eldest_kid"`
			EldestKeyFingerprint string        `json:"eldest_key_fingerprint"`
			PgpPublicKeys        []interface{} `json:"pgp_public_keys"`
		} `json:"public_keys"`
		ProofsSummary struct {
			ByPresentationGroup struct {
			} `json:"by_presentation_group"`
			BySigID struct {
			} `json:"by_sig_id"`
			All    []interface{} `json:"all"`
			HasWeb bool          `json:"has_web"`
		} `json:"proofs_summary"`
		CryptocurrencyAddresses struct {
		} `json:"cryptocurrency_addresses"`
		Sigs struct {
			Last struct {
				SigID       string `json:"sig_id"`
				Seqno       int    `json:"seqno"`
				PayloadHash string `json:"payload_hash"`
			} `json:"last"`
		} `json:"sigs"`
		Devices struct {
			FourD3A722B4D570A437Cca9Cf4B05A0518 struct {
				Type   string `json:"type"`
				Ctime  int    `json:"ctime"`
				Mtime  int    `json:"mtime"`
				Name   string `json:"name"`
				Status int    `json:"status"`
				Keys   []struct {
					Kid     string `json:"kid"`
					KeyRole int    `json:"key_role"`
					SigID   string `json:"sig_id"`
				} `json:"keys"`
			} `json:"4d3a722b4d570a437cca9cf4b05a0518"`
		} `json:"devices"`
		Stellar struct {
			Hidden  bool `json:"hidden"`
			Primary struct {
			} `json:"primary"`
		} `json:"stellar"`
	} `json:"them"`
}

func keybaseExtra(userres string) string {
	var resp = http.GetReq("https://keybase.io/_/api/1.0/user/lookup.json?usernames=" + userres)
	var s = new(keybase)
	err := json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", s.Them[0].ID, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Username: ", s.Them[0].Basics.Username, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ cTime: ", strconv.Itoa(s.Them[0].Basics.Ctime), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ mTime: ", strconv.Itoa(s.Them[0].Basics.Mtime), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID Version: ", strconv.Itoa(s.Them[0].Basics.IDVersion), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Track Version: ", strconv.Itoa(s.Them[0].Basics.TrackVersion), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Last ID Change: ", strconv.Itoa(s.Them[0].Basics.LastIDChange), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Username Cased: ", s.Them[0].Basics.UsernameCased, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Status: ", strconv.Itoa(s.Them[0].Basics.Status), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Salt: ", s.Them[0].Basics.Salt, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Eldest Seqno: ", strconv.Itoa(s.Them[0].Basics.EldestSeqno), firstSuccess)
	return r
}

type issuu struct {
	Rsp struct {
		Content struct {
			User struct {
				ID                     int    `json:"id"`
				Username               string `json:"username"`
				Password               bool   `json:"password"`
				Account                string `json:"account"`
				Displayname            string `json:"displayName"`
				Country                string `json:"country"`
				Stackcount             int    `json:"stackCount"`
				Stackcountpublic       int    `json:"stackCountPublic"`
				Documentcount          int    `json:"documentCount"`
				Subscribercount        int    `json:"subscriberCount"`
				Subscriptioncount      int    `json:"subscriptionCount"`
				Likeddocumentcount     int    `json:"likedDocumentCount"`
				Clippingcount          int    `json:"clippingCount"`
				Stacksubscriptioncount int    `json:"stackSubscriptionCount"`
			} `json:"user"`
		} `json:"_content"`
		Stat string `json:"stat"`
	} `json:"rsp"`
}

func issuuExtra(userres string) string {
	var resp = http.GetReq("https://issuu.com/query?format=json&_=3210224608766&profileUsername=" + userres + "&action=issuu.user.get_anonymous")
	var s = new(issuu)
	err := json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", strconv.Itoa(s.Rsp.Content.User.ID), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Username: ", s.Rsp.Content.User.Username, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Password: ", strconv.FormatBool(s.Rsp.Content.User.Password), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Account: ", s.Rsp.Content.User.Account, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Display Name: ", s.Rsp.Content.User.Displayname, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Country: ", s.Rsp.Content.User.Country, firstSuccess)
	return r
}

type disqus struct {
	Code     int `json:"code"`
	Response struct {
		Disable3Rdpartytrackers bool    `json:"disable3rdPartyTrackers"`
		Ispowercontributor      bool    `json:"isPowerContributor"`
		Isprimary               bool    `json:"isPrimary"`
		ID                      string  `json:"id"`
		Numfollowers            int     `json:"numFollowers"`
		Rep                     float64 `json:"rep"`
		Numfollowing            int     `json:"numFollowing"`
		Numposts                int     `json:"numPosts"`
		Location                string  `json:"location"`
		Isprivate               bool    `json:"isPrivate"`
		Joinedat                string  `json:"joinedAt"`
		Username                string  `json:"username"`
		Numlikesreceived        int     `json:"numLikesReceived"`
		Reputationlabel         string  `json:"reputationLabel"`
		About                   string  `json:"about"`
		Name                    string  `json:"name"`
		URL                     string  `json:"url"`
		Numforumsfollowing      int     `json:"numForumsFollowing"`
		Profileurl              string  `json:"profileUrl"`
		Reputation              float64 `json:"reputation"`
		Avatar                  struct {
			Small struct {
				Permalink string `json:"permalink"`
				Cache     string `json:"cache"`
			} `json:"small"`
			Iscustom  bool   `json:"isCustom"`
			Permalink string `json:"permalink"`
			Cache     string `json:"cache"`
			Large     struct {
				Permalink string `json:"permalink"`
				Cache     string `json:"cache"`
			} `json:"large"`
		} `json:"avatar"`
		Signedurl   string `json:"signedUrl"`
		Isanonymous bool   `json:"isAnonymous"`
	} `json:"response"`
}

func disqusExtra(userres string) string {
	var resp = http.GetReq("https://disqus.com/api/3.0/users/details?user=username%3A" + userres + "&attach=userFlaggedUser&api_key=E8Uh5l5fHZ6gD8U3KycjAIAk46f68Zw7C6eW8WSjZvCLXebZ7p0r1yrYDrLilk2F")
	var s = new(disqus)
	err := json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", s.Response.ID, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Location: ", s.Response.Location, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ About: ", s.Response.About, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Name: ", s.Response.Name, firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Private: ", strconv.FormatBool(s.Response.Isprivate), firstSuccess)
	r, firstSuccess = cli.TreeIt(r, "      ├─ Anonymous: ", strconv.FormatBool(s.Response.Isanonymous), firstSuccess)
	return r
}

/*

TODO

*/

type gitlab []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

func gitlabExtra(userres string) string {
	var resp = http.GetReq("https://gitlab.com/api/v4/users?username=" + userres)
	var s = new(gitlab)
	err := json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return ""
	}
	var r string
	firstSuccess := false
	for _, value := range *s {
		r, firstSuccess = cli.TreeIt(r, "      ├─ ID: ", strconv.Itoa(value.ID), firstSuccess)
		r, firstSuccess = cli.TreeIt(r, "      ├─ Name: ", value.Name, firstSuccess)
		r, firstSuccess = cli.TreeIt(r, "      ├─ State: ", value.State, firstSuccess)
		r, firstSuccess = cli.TreeIt(r, "      ├─ Avatar URL: ", value.AvatarURL, firstSuccess)
		r, firstSuccess = cli.TreeIt(r, "      ├─ Username: ", value.Username, firstSuccess)
		r, firstSuccess = cli.TreeIt(r, "      ├─ Web URL: ", value.WebURL, firstSuccess)
	}

	return r
}
