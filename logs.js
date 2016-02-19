window.addEventListener('load', function () {
    (function($) {
	$.QueryString = (function(a) {
            if (a == "") return {};
            var b = {};
            for (var i = 0; i < a.length; ++i)
            {
		var p=a[i].split('=');
		if (p.length != 2) continue;
		b[p[0]] = decodeURIComponent(p[1].replace(/\+/g, " "));
            }
            return b;
	})(window.location.search.substr(1).split('&'))
    })(jQuery);
    
    var url = 'ws://' + window.location.hostname + ':' + window.location.port + '/ws?addr='+$.QueryString["addr"]
    console.log("Connecting to "+url);
    var ws = new WebSocket(url);

    ws.onclose = function(close) {
	//console.log(close);
    }

    ws.onerror = function(error){
	//console.log(error);
    }

    ws.onopen = function() {
	$("ul").append("Connected!");
	
	ws.onmessage = function(msg) {
	    var evt = JSON.parse(msg.data)
	    console.log(evt)

	    switch (evt.event) {
	    case "PlayerConnected":
		$("ul").append("<li> Player "+evt.data.username+" has connected </li>")
		break;
	    case "PlayerDisconnected":
		$("ul").append("<li> Player "+evt.data.username+" has disconnected </li>")
		break;
	    case "PlayerGlobalMessage":
		$("ul").append("<li> "+evt.data.username+": <b>"+evt.data.text+"</b></li>")
		break;
	    case "PlayerTeamMessage":
		$("ul").append("<li> "+evt.data.username+": <b> (team)"+evt.data.text+"</b></li>")
		break;
	    case "PlayerSpawned":
		$("ul").append("<li> "+evt.data.username+" spawned as <b>"+evt.data.class+"</b></li>")
		break;
	    case "PlayerClassChange":
		$("ul").append("<li> "+evt.data.username+" changed class to <b>"+evt.data.class+"</b></li>")
		break;
	    case "PlayerKilled":
		$("ul").append("<li> "+evt.data.player1.username+" killed "+evt.data.player2.username+" with "+evt.data.weapon+" "+evt.data.customkill+"</b></li>")
		break;
	    case "PlayerDamaged":
		if (evt.data.player1.damage <= 10) {
		    break;
		}
		
		var str = "<li> "+evt.data.player1.username+ " <b>damaged</b> "+evt.data.player2.username+" with <b>"+evt.data.weapon+"</b> for <font color='red'>"+evt.data.damage+"</font> health ";
		if (evt.data.airshot) {
		    str += "<b>(AIRSHOT)</b>"
		}
		if (evt.data.headshot) {
		    str += "<b>(HEADSHOT)</b>"
		}
		str += "</li>"
		$("ul").append(str)

		break;
	    case "PlayerHealed":
		var str = "<li> "+evt.data.player1.username+" <b>healed</b> "+evt.data.player2.username+" for <font color='red'>"+evt.data.healed+"</font> health</li>"
		$("ul").append(str)
	    case "PlayerKilledMedic"
		$("ul").append("<li> "+evt.data.player1.username+" <b>killed medic</b> "+evt.data.player2.username+"</li>")
	    case "PlayerUberFinished":
		$("ul").append("<li> <b>ubercharge finished</b> for "+evt.data.player1.username+"</li>")
	    case "PlayerBlockedCapture":
		$("ul").append("<li> capture of point "+evt.data.cp+", "+evt.data.cpname+" <b>blocked</b> by "+evt.data.player1.username)
	    case "PlayerItemPickup":
		var str = "<li> "+evt.data.player.username+" picked up "+evt.data.item
		if evt.data.healing != 0 {
		    str += " (<b>healing "+evt.data.healing+"</b>) </li>"
		}

		$("ul").append(str)
	    case "GameOver":
		$("ul").append("<li> Game over </li>")
	    case "WorldRoundWin":
		$("ul").append("<li> Team "+evt.data.team+" Won </li>")
	    case "TeamScoreUpdate":
		$("ul").append("<li> Team "+evt.data.team+" has "+evt.data.score+" points</li>")
	    case "WorldRoundWin":
		$("ul").append("<li> Team"+evt.data.team+" won </li>")
	    }
	}
    }
}
		       );
