let Print = props =>{
    console.log(props);
}

let TestNav = props =>{
    console.log(props.item)
    var data = props.item
    const listItems = data.map((item) =>
        <li key={item}>{item}</li>
    )
    return <ul>{listItems}</ul>
}
let nav = (data) => {
    console.log(data)
    ReactDOM.render(
        <TestNav item={data} />,
        document.getElementById('test')
    )
}


let AuthorContent = (props) => {
    var data = props.items
    return <ul>
        <li>{data.ID}</li>
        <li>{data.name}</li>
        <li>{data.email}</li>
        <li>{data.url}</li>
    </ul>
    
}
let author = (data) => {
    console.log(data)
    ReactDOM.render(
        <AuthorContent items={data} />,
        document.getElementById('test')
    )
}