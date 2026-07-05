import React from "react";

const MakeRoom = ()=>{
    const makeRoom =  async (formData) =>{


        var roomName = formData.get("roomName");
        
        if (roomName) {
           //Make room
           const response = await fetch("http://localhost:8080/makeRoom", {
                method: "POST",
                body : JSON.stringify({
                    "roomName" : roomName
                }) 
           })
    
        }
        else {
            console.log("No name given")
            
        }
    }
    return (
        <form action={makeRoom}>
            <label> Room Name </label>
            <input type="text" name="roomName"></input>
            <button type="submit">Submit</button>
        </form>
    )
}

export default MakeRoom;