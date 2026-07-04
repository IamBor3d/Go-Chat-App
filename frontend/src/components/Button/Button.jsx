

const Button = ({text, onClick}) =>{
    return (
        <div className="border radius-xs border-gray-200 p-4 m-3">
            <button onClick={onClick}>{text}</button>
        </div>
    )
}

export default Button;