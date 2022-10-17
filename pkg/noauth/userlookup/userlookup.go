package userlookup

import (
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Website is used as the response type of the Se
type Website struct {
	Title  string
	Domain string
	Delete string
	Valid  bool
	Extra  string
}

var results []Website

// UserLookup returns a slice of type Website
func UserLookup(userres string) []Website {
	var wg sync.WaitGroup
	var arrNo = noRedirSites(userres)
	for _, v := range arrNo {
		go checkUser(v, false, &wg)
	}
	var arrYes = redirSites(userres)
	for _, v := range arrYes {
		go checkUser(v, true, &wg)
	}

	wg.Wait()
	return results
}

func noRedirSites(userres string) []Website {
	var arr []Website

	arr = append(arr, Website{Title: "FACEBOOK", Domain: "https://www.facebook.com/" + userres, Delete: "https://www.facebook.com/help/delete_account?rdrhc"}, Website{Title: "BLOGSPOT", Domain: "https://" + userres + ".blogspot.com", Delete: "TODO"}, Website{Title: "WORDPRESS", Domain: "https://" + userres + ".wordpress.com", Delete: "N/A"})

	return arr
}

func redirSites(userres string) []Website {
	var arr []Website

	arr = append(arr, Website{Title: "PINTEREST", Domain: "https://www.pinterest.com/" + userres, Delete: "https://pinterest.com/settings/"}, Website{Title: "GITHUB", Domain: "https://www.github.com/" + userres, Delete: "https://github.com/settings/admin"}, Website{Title: "TUMBLR", Domain: "https://" + userres + ".tumblr.com", Delete: "https://www.tumblr.com/account/delete"}, Website{Title: "FLICKR", Domain: "https://www.flickr.com/people/" + userres, Delete: "http://www.flickr.com/profile_delete.gne"}, Website{Title: "VIMEO", Domain: "https://vimeo.com/" + userres, Delete: "https://vimeo.com/settings/goodbye/forever"}, Website{Title: "SOUNDCLOUD", Domain: "https://soundcloud.com/" + userres, Delete: "http://soundcloud.com/settings/account#delete-user"}, Website{Title: "DISQUS", Domain: "https://disqus.com/by/" + userres, Delete: "http://disqus.com/pages/dashboard/#account"}, Website{Title: "MEDIUM", Domain: "https://medium.com/@" + userres, Delete: "https://medium.com/me/settings"}, Website{Title: "DEVIANTART", Domain: "https://" + userres + ".deviantart.com", Delete: "https://www.deviantart.com/settings/deactivation"}, Website{Title: "VK", Domain: "https://vk.com/" + userres, Delete: "http://vk.com/settings?act=deactivate"}, Website{Title: "ABOUTME", Domain: "https://about.me/" + userres, Delete: "https://about.me/account"}, Website{Title: "FLIPBOARD", Domain: "https://flipboard.com/@" + userres, Delete: "https://flipboard.com/support/"}, Website{Title: "SLIDESHARE", Domain: "https://slideshare.net/" + userres, Delete: "N/A"}, Website{Title: "SCRIBD", Domain: "https://www.scribd.com/" + userres, Delete: "http://www.scribd.com/account_settings/preferences"}, Website{Title: "PATREON", Domain: "https://www.patreon.com/" + userres, Delete: "TODO"}, Website{Title: "BITBUCKET", Domain: "https://bitbucket.org/" + userres, Delete: "https://bitbucket.org/account/"}, Website{Title: "DAILYMOTION", Domain: "https://www.dailymotion.com/" + userres, Delete: "TODO"}, Website{Title: "ETSY", Domain: "https://www.etsy.com/shop/" + userres, Delete: "http://www.etsy.com/help/article/53"}, Website{Title: "CASHAPP", Domain: "https://cash.app/$" + userres, Delete: "TODO"}, Website{Title: "BEHANCE", Domain: "https://www.behance.net/" + userres, Delete: "http://www.behance.net/account/privacy"}, Website{Title: "GOODREADS", Domain: "https://www.goodreads.com/" + userres, Delete: "http://www.goodreads.com/user/destroy"}, Website{Title: "INSTRUCTABLES", Domain: "https://www.instructables.com/member/" + userres, Delete: "Email service@instructables.com"}, Website{Title: "KEYBASE", Domain: "https://keybase.io/" + userres, Delete: "TODO"}, Website{Title: "KONGREGATE", Domain: "https://kongregate.com/accounts/" + userres, Delete: "http://www.kongregate.com/forums/7/topics/241772?page=1#posts-5207538"}, Website{Title: "LIVEJOURNAL", Domain: "https://" + userres + ".livejournal.com", Delete: "http://www.livejournal.com/accountstatus.bml"}, Website{Title: "LASTFM", Domain: "https://last.fm/user/" + userres, Delete: "http://www.last.fm/settings/account"}, Website{Title: "DRIBBBLE", Domain: "https://dribbble.com/" + userres, Delete: "http://dribbble.com/account"}, Website{Title: "CODECADEMY", Domain: "https://www.codecademy.com/profiles/" + userres, Delete: "http://help.codecademy.com/customer/portal/articles/1394081-how-do-i-delete-my-account-"}, Website{Title: "GRAVATAR", Domain: "https://en.gravatar.com/" + userres, Delete: "N/A"}, Website{Title: "PASTEBIN", Domain: "https://pastebin.com/u/" + userres, Delete: "N/A"}, Website{Title: "ROBLOX", Domain: "https://www.roblox.com/user.aspx?username=" + userres, Delete: "N/A"}, Website{Title: "GUMROAD", Domain: "https://www.gumroad.com/" + userres, Delete: "https://gumroad.com/settings"}, Website{Title: "NEWGROUNDS", Domain: "https://" + userres + ".newgrounds.com", Delete: "TODO"}, Website{Title: "WATTPAD", Domain: "https://www.wattpad.com/user/" + userres, Delete: "TODO"}, Website{Title: "TRAKT", Domain: "https://www.trakt.tv/users/" + userres, Delete: "http://trakt.tv/settings/data"}, Website{Title: "BUZZFEED", Domain: "https://buzzfeed.com/" + userres, Delete: "TODO"}, Website{Title: "TRIPADVISOR", Domain: "https://www.tripadvisor.com/Profile/" + userres, Delete: "https://www.tripadvisorsupport.com/hc/en-us/articles/200615117"}, Website{Title: "HUBPAGES", Domain: "https://" + userres + ".hubpages.com", Delete: "TODO"}, Website{Title: "BLIPFM", Domain: "https://blip.fm/" + userres, Delete: "TODO"}, Website{Title: "WIKIPEDIA", Domain: "https://en.wikipedia.org/wiki/User:" + userres, Delete: "https://en.wikipedia.org/wiki/Wikipedia:FAQ#How_do_I_change_my_username.2Fdelete_my_account.3F"}, Website{Title: "CODEMENTOR", Domain: "https://www.codementor.io/" + userres, Delete: "TODO"}, Website{Title: "REVERBNATION", Domain: "https://www.reverbnation.com/" + userres, Delete: "TODO"}, Website{Title: "DESIGNSPIRATION", Domain: "https://www.designspiration.com/" + userres, Delete: "TODO"}, Website{Title: "BANDCAMP", Domain: "https://www.bandcamp.com/" + userres, Delete: "https://bandcamp.com/settings?pane=fan"}, Website{Title: "COLOURLOVERS", Domain: "https://www.colourlovers.com/lover/" + userres, Delete: "TODO"}, Website{Title: "IFTTT", Domain: "https://www.ifttt.com/p/" + userres, Delete: "https://ifttt.com/settings/deactivate"}, Website{Title: "SLACK", Domain: "https://" + userres + ".slack.com", Delete: "https://my.slack.com/account/settings"}, Website{Title: "ELLO", Domain: "https://ello.co/" + userres, Delete: "https://ello.co/delete-account"}, Website{Title: "7CUPS", Domain: "https://www.7cups.com/@" + userres + "", Delete: "TODO"}, Website{Title: "INDEPENDENT ACADEMIA", Domain: "https://independacademia.edu/" + userres + "", Delete: "TODO"}, Website{Title: "ALIK", Domain: "https://www.alik.cz/u/" + userres + "", Delete: "TODO"}, Website{Title: "AUDIO JUNGLE", Domain: "https://audiojungle.net/user/" + userres + "", Delete: "https://help.market.envato.com/hc/en-us/articles/202500394-How-Do-I-Close-My-Account-"}, Website{Title: "BOOK CROSSING", Domain: "https://www.bookcrossing.com/mybookshelf/" + userres + "/", Delete: "TODO"}, Website{Title: "BUY ME A COFFEE", Domain: "https://buymeacoff.ee/" + userres + "", Delete: "TODO"}, Website{Title: "CNET", Domain: "https://www.cnet.com/profiles/" + userres + "/", Delete: "https://cbsi.secure.force.com/CBSi/submitcase?template=template_cnet&referer=cnet.com&cfs=SFS_1"}, Website{Title: "CARBONMADE", Domain: "https://" + userres + ".carbonmade.com", Delete: "TODO"}, Website{Title: "Chess", Domain: "https://www.chess.com/member/" + userres + "", Delete: "TODO"}, Website{Title: "CodeWars", Domain: "https://www.codewars.com/users/" + userres + "", Delete: "TODO"}, Website{Title: "Countable", Domain: "https://www.countable.us/" + userres + "", Delete: "TODO"}, Website{Title: "DevTo", Domain: "https://dev.to/" + userres + "", Delete: "TODO"}, Website{Title: "Discogs", Domain: "https://www.discogs.com/user/" + userres + "", Delete: "http://www.discogs.com/users/delete"}, Website{Title: "Eyeem", Domain: "https://www.eyeem.com/u/" + userres + "", Delete: "TODO"}, Website{Title: "F3", Domain: "https://f3.cool/" + userres + "/", Delete: "TODO"}, Website{Title: "Fortnite Tracker", Domain: "https://fortnitetracker.com/profile/all/" + userres + "", Delete: "TODO"}, Website{Title: "FreeSound", Domain: "https://freesound.org/people/" + userres + "/", Delete: "http://www.freesound.org/home/delete/"}, Website{Title: "GameSpot", Domain: "https://www.gamespot.com/profile/" + userres + "/", Delete: "TODO"}, Website{Title: "Giphy", Domain: "https://giphy.com/" + userres + "", Delete: "TODO"}, Website{Title: "Gitlab", Domain: "https://gitlab.com/" + userres + "", Delete: "TODO"}, Website{Title: "GuruShots", Domain: "https://gurushots.com/" + userres + "/photos", Delete: "TODO"}, Website{Title: "Hackaday", Domain: "https://hackaday.io/" + userres + "", Delete: "TODO"}, Website{Title: "HackerOne", Domain: "https://hackerone.com/" + userres + "", Delete: "TODO"}, Website{Title: "Houzz", Domain: "https://houzz.com/user/" + userres + "", Delete: "TODO"}, Website{Title: "Issuu", Domain: "https://issuu.com/" + userres + "", Delete: "https://issuu.com/home/settings/billing"}, Website{Title: "Itch", Domain: "https://" + userres + ".itch.io/", Delete: "TODO"}, Website{Title: "Jimdosite", Domain: "https://" + userres + ".jimdosite.com", Delete: "TODO"}, Website{Title: "LeetCode", Domain: "https://leetcode.com/" + userres + "", Delete: "TODO"}, Website{Title: "Letterboxd", Domain: "https://letterboxd.com/" + userres + "", Delete: "http://letterboxd.com/user/disableaccount/"}, Website{Title: "Lichess", Domain: "https://lichess.org/@/" + userres + "", Delete: "TODO"}, Website{Title: "Memrise", Domain: "https://www.memrise.com/user/" + userres + "/", Delete: "http://www.memrise.com/settings/deactivate/"}, Website{Title: "Munzee", Domain: "https://www.munzee.com/m/" + userres + "", Delete: "TODO"}, Website{Title: "My Anime List", Domain: "https://myanimelist.net/profile/" + userres + "", Delete: "TODO"}, Website{Title: "My Mini Factory", Domain: "https://www.myminifactory.com/users/" + userres + "", Delete: "TODO"}, Website{Title: "MySpace", Domain: "https://myspace.com/" + userres + "", Delete: "https://myspace.com/settings/profile"}, Website{Title: "Periscope", Domain: "https://www.periscope.tv/" + userres + "/", Delete: "TODO"}, Website{Title: "Pink Bike", Domain: "https://www.pinkbike.com/u/" + userres + "/", Delete: "TODO"}, Website{Title: "Pokemon Showdown", Domain: "https://pokemonshowdown.com/users/" + userres + "", Delete: "TODO"}, Website{Title: "Product Hunt", Domain: "https://www.producthunt.com/@" + userres + "", Delete: "TODO"}, Website{Title: "PromoDJ", Domain: "http://promodj.com/" + userres + "", Delete: "TODO"}, Website{Title: "PyPi", Domain: "https://pypi.org/user/" + userres + "", Delete: "TODO"}, Website{Title: "Rajce", Domain: "https://" + userres + ".rajce.idnes.cz/", Delete: "TODO"}, Website{Title: "Rate Your Music", Domain: "https://rateyourmusic.com/~" + userres + "", Delete: "http://rateyourmusic.com/account/delete"}, Website{Title: "Red Bubble", Domain: "https://www.redbubble.com/people/" + userres + "", Delete: "TODO"}, Website{Title: "Repl.it", Domain: "https://repl.it/@" + userres + "", Delete: "TODO"}, Website{Title: "Sbazar", Domain: "https://www.sbazar.cz/" + userres + "", Delete: "TODO"}, Website{Title: "Shitpost Bot", Domain: "https://www.shitpostbot.com/user/" + userres + "", Delete: "TODO"}, Website{Title: "Smule", Domain: "https://www.smule.com/" + userres + "", Delete: "TODO"}, Website{Title: "SourceForge", Domain: "https://sourceforge.net/u/" + userres + "", Delete: "https://sourceforge.net/auth/disable/"}, Website{Title: "Speedrun", Domain: "https://speedrun.com/user/" + userres + "", Delete: "TODO"}, Website{Title: "Roberts Space Industries", Domain: "https://robertsspaceindustries.com/citizens/" + userres + "", Delete: "TODO"}, Website{Title: "Tellonym", Domain: "https://tellonym.me/" + userres + "", Delete: "TODO"}, Website{Title: "Trading View", Domain: "https://www.tradingview.com/u/" + userres + "/", Delete: "TODO"}, Website{Title: "Ultimate Guitar", Domain: "https://ultimate-guitar.com/u/" + userres + "", Delete: "TODO"}, Website{Title: "Unsplash", Domain: "https://unsplash.com/@" + userres + "", Delete: "TODO"}, Website{Title: "VSCO", Domain: "https://vsco.co/" + userres + "", Delete: "TODO"}, Website{Title: "Vero", Domain: "https://vero.co/" + userres + "", Delete: "TODO"}, Website{Title: "WarriorForum", Domain: "https://www.warriorforum.com/members/" + userres + ".html", Delete: "TODO"}, Website{Title: "We Heart It", Domain: "https://weheartit.com/" + userres + "", Delete: "TODO"}, Website{Title: "Windy", Domain: "https://community.windy.com/user/" + userres + "", Delete: "TODO"}, Website{Title: "XBox", Domain: "https://xboxgamertag.com/search/" + userres + "", Delete: "TODO"}, Website{Title: "Zhihu", Domain: "https://www.zhihu.com/people/" + userres + "", Delete: "TODO"}, Website{Title: "All My Links", Domain: "https://allmylinks.com/" + userres + "", Delete: "TODO"}, Website{Title: "Couch Surfing", Domain: "https://www.couchsurfing.com/people/" + userres + "", Delete: "https://www.couchsurfing.org/delete_profile.html?delete=1"}, Website{Title: "DailyKOs", Domain: "https://www.dailykos.com/user/" + userres + "", Delete: "TODO"}, Website{Title: "Dating", Domain: "http://dating.ru/" + userres + "", Delete: "TODO"}, Website{Title: "Drive2", Domain: "https://www.drive2.ru/users/" + userres + "", Delete: "TODO"}, Website{Title: "FL", Domain: "https://www.fl.ru/users/" + userres + "", Delete: "TODO"}, Website{Title: "Geocaching", Domain: "https://www.geocaching.com/p/default.aspx?u=" + userres + "", Delete: "http://support.groundspeak.com/index.php?pg=kb.page&id=102"}, Website{Title: "Gfycat", Domain: "https://gfycat.com/@" + userres + "", Delete: "TODO"}, Website{Title: "Hackster", Domain: "https://www.hackster.io/" + userres + "", Delete: "TODO"}, Website{Title: "Irecommend", Domain: "https://irecommend.ru/users/" + userres + "", Delete: "TODO"}, Website{Title: "JBZD", Domain: "https://jbzd.com.pl/uzytkownik/" + userres + "", Delete: "TODO"}, Website{Title: "JeuxVideo", Domain: "http://www.jeuxvideo.com/profil/" + userres + "?mode=infos", Delete: "TODO"}, Website{Title: "Kwork", Domain: "https://kwork.ru/user/" + userres + "", Delete: "TODO"}, Website{Title: "Livelib", Domain: "https://www.livelib.ru/reader/" + userres + "", Delete: "TODO"}, Website{Title: "MOIKRUG", Domain: "https://moikrug.ru/" + userres + "", Delete: "TODO"}, Website{Title: "NAIRALAND", Domain: "https://www.nairaland.com/" + userres + "", Delete: "TODO"}, Website{Title: "NN", Domain: "https://" + userres + ".www.nn.ru/", Delete: "TODO"}, Website{Title: "NOTE", Domain: "https://note.com/" + userres + "", Delete: "TODO"}, Website{Title: "NPMJS", Domain: "https://www.npmjs.com/~" + userres + "", Delete: "TODO"}, Website{Title: "OSU", Domain: "https://osu.ppy.sh/users/" + userres + "", Delete: "TODO"}, Website{Title: "PIKABU", Domain: "https://pikabu.ru/@" + userres + "", Delete: "TODO"}, Website{Title: "TOSTER", Domain: "https://www.toster.ru/user/" + userres + "/answers", Delete: "TODO"}, Website{Title: "PORNHUB", Domain: "https://pornhub.com/users/" + userres}, Website{Title: "XVIDEOS", Domain: "https://xvideos.com/profiles/" + userres}, Website{Title: "YOUPORN", Domain: "https://youporn.com/uservids/" + userres, Delete: "TODO"}, Website{Title: "XHAMSTER", Domain: "https://xhamster.com/users/" + userres, Delete: "TODO"}, Website{Title: "BONGACAMS", Domain: "https://pt.bongacams.com/profile/" + userres, Delete: "TODO"}, Website{Title: "ROYALCAMS", Domain: "https://royalcams.com/profile/" + userres, Delete: "TODO"}, Website{Title: "CHATURBATE", Domain: "https://chaturbate.com/" + userres, Delete: "TODO"})

	return arr
}

func getSCredir(title, url string, redirect bool) Website {

	method := "GET"
	var client *http.Client
	if redirect {
		client = &http.Client{
			Timeout: 5 * time.Second,
		}
	} else {
		client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 5 * time.Second,
		}
	}
	req, err := http.NewRequest(method, url, http.NoBody)

	if err != nil {
		return Website{Title: title, Domain: url, Valid: false}
	}
	res, err := client.Do(req)
	if err != nil {
		return Website{Title: title, Domain: url, Valid: false}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Website{Title: title, Domain: url, Valid: false}
	}
	strconv.AppendBool(body, true)
	if res.StatusCode == 200 {
		return Website{Title: title, Domain: url, Valid: true}
	}
	return Website{Title: title, Domain: url, Valid: false}

}

func checkUser(site Website, redirect bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	x := getSCredir(site.Title, site.Domain, redirect)
	results = append(results, x)
}
