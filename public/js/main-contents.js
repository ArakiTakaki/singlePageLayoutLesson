export default nav = (data) => {
    console.log(data)
    ReactDOM.render(
        <TestNav item={data} />,
        document.getElementById('test')
    )
}

let TestNav = props =>{
    console.log(props.item)
    var data = props.item
    const listItems = data.map((item) =>
        <li key={item}>{item}</li>
    )
    return <ul>{listItems}</ul>
}

