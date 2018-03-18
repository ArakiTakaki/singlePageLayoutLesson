
let test = new Transition();
test.comments()

let tests = (data) => {
    console.log(data)
}


let TestNav = props =>{
    console.log(props.item)
    var data = props.item
    const listItems = data.map((item) =>
        <li key={item}>{item}</li>
    )
    return <ul>{listItems}</ul>
}
let aaa = (data) => {
    console.log(data)
    ReactDOM.render(
        <TestNav item={data} />,
        document.getElementById('test')
    )
}

test.getJson("/api/test",aaa)