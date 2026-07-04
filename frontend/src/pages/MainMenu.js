import React, { useRef } from "react";
import { Link } from "react-router-dom";
import Button from "../components/Button/Button";

const MainMenu = () => {

    return (
        <div className=" flex m-12 p-3 text-center justify-center">

            <div className="grid grid-rows-1 grid-cols-2 gap-4 content-center">
                <Link to="/makeRoom"> <Button text="Make Room"/> </Link>
                <Link to="/rooms"> <Button text={"Join Rooms"}/></Link>
            </div>

        </div>
        
    )
}

export default MainMenu;