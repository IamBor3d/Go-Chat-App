const Message = ({text}) => {

    let temp = JSON.parse(text);
    console.log(temp)

    return (
        <div className="p-2">
            <div className="shadow-xl border rounded-sm border-gray-100 p-10px">
            {temp.Body}
            </div>
        </div>
        );
        

};

export default Message;