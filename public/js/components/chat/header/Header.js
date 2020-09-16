import React, { Component } from "react"
import VHeader from './VHeader'

class Header extends Component{

    constructor(props){
        super(props)
        this.view = VHeader()
    }

    render(){
        return this.view
    }
}

export default Header