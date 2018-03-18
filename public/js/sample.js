
class Transition {

    comments(){
        console.log("test")
    }
    // trans JsonのURLとwork()メソッドを格納  返り値はJSONで帰ってくる。
    getJson(url,work){
        var myHeaders = new Headers()
        var myInit = {
            method: 'POST',
            headers: myHeaders,
            body:"",
            mode: 'cors',
            catch: 'default'
        };
        var myRequest = new Request(url, myInit);
        fetch(myRequest)
            .then(res => res.json())
            .then(data => work(data) )
            .catch(e => console.error('Error:', e));        
    }
}