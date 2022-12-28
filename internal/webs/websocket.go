package webs

import (
	"flag"
	"os"
	"strconv"
	"sync"

	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8484", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request, al *ArkLog) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error upgrade:", err)
		return
	}
	// засунем или перепишем в MAP
	al.conn[c] = true
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v, user-agent: %v", err, r.Header.Get("User-Agent"))
			}
			delete(al.conn, c)
			log.Println("read:", err)
			break
		}
		// log - мой лог перехватывает сообщение
		log.Printf("с сайта: %s", message)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

type ArkLog struct {
	wspool
}
type wspool struct {
	// Registered connections. That's a connection pool
	conn map[*websocket.Conn]bool
	mu   sync.Mutex
}

func (ws *wspool) sendMess(wg *sync.WaitGroup, c *websocket.Conn, mess []byte) {
	defer wg.Done()
	ws.mu.Lock()
	defer ws.mu.Unlock()

	ppp := "[" + strconv.Itoa(len(ws.conn)) + "] " // + ppp
	ppp = ppp + ": "
	ppp1 := append([]byte(ppp), mess...)
	err := c.WriteMessage(websocket.TextMessage, ppp1)
	if err != nil {
		log.Println("Ошибка отправки в сокет", c.RemoteAddr().String())
		delete(ws.conn, c)
		if err := c.Close(); err != nil {
			log.Println("Ошибка закрытия сокета")
		}
	}
}

func (a *ArkLog) Write(p []byte) (n int, err error) {
	// вывод на консоль
	k, e := os.Stdout.Write(p)
	// вывод в сокет
	var wg sync.WaitGroup
	//	go func() {
	for c := range a.conn {
		wg.Add(1)
		go a.sendMess(&wg, c, p)
	}
	//	}()
	wg.Wait()
	return k, e
}

func Start() {
	arklog := &ArkLog{}
	arklog.conn = make(map[*websocket.Conn]bool)
	log.SetOutput(arklog)

	flag.Parse()

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		echo(w, r, arklog)
	})
	http.HandleFunc("/", home)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Println("WS Server Error:", err.Error())
	}

}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
	var count = 0;
    var print = function(message) {
		var numberOfChildren = output.getElementsByTagName('div').length;
        var d = document.createElement("div");
        if (numberOfChildren && numberOfChildren > 300) {
			var removableNode = output.querySelectorAll("div")[0];
		   // удаляем узел
		   output.removeChild(removableNode);
		}
		count++;
        d.textContent =  count+ " > "+message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    
    
	let wsconnect = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("Подключилися");
        }
        ws.onclose = function(evt) {
            print("Отключен, ждемс..."); // срабатывает после ошибки тоже
			if ( ws ) {
               ws = null;
			}
			setTimeout(function() {
				wsconnect();
			  }, 1000)
        }
        ws.onmessage = function(evt) {
            print("" + evt.data);
        }
        ws.onerror = function(evt) {
			
            print("вот такая ошибка: " + evt.data);
			ws.close()
        }
        return false;
    };
	wsconnect()
	// document.getElementById("open").onclick = wsconnect

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("Туда: " + input.value);
        ws.send(input.value);
        return false;
    };
    // document.getElementById("close").onclick = function(evt) {
    //     if (!ws) {
    //         return false;
    //     }
    //     ws.close();
    //     return false;
    // };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="10%">
<!-- <p>Открыть - подключаемся</p></ br>
<p> Закрыть - отключаемся</p></ by> 
<p>Отправить - Отправляем</p></ br> 
-->

<form>
<!-- <button id="open">Открыть</button>
<button id="close">Закрыть</button> -->
<p><input id="input" type="text" value="Привет!">
<button id="send">Отправить</button>
</form>
</td><td valign="top" width="90%">
<div id="output" style="max-height: 90vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
