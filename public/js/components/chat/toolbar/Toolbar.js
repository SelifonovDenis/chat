import React, { Component } from "react"
import VToolbar from './VToolbar'
import Request from "../../providers/Request";

class Toolbar extends Component{

    constructor(props){
        super(props)
        this.view = VToolbar.bind(this)
        this.state = {
            logout: false,
        }
        this.request = new Request()
    }

    logout = () => {
        this.request.Post('logout', this.props.user)
            .then((res)=>{
                if (res.type === "data") {
                    this.setState(()=>({
                        logout: true
                    }))
                } else{
                    console.log(res)
                }
            }).catch((err)=>{
            console.log(err)
        })
    }

    render(){
        console.log(this.props.user)
        if (!this.props.user){
            return <div></div>
        }
        return this.view()
    }
}

export default Toolbar