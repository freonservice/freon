import React from "react";
import {Link} from "react-router-dom";
import {
    Navbar,
    Nav,
    NavItem,
} from "./../../components";

import {NavbarUser} from "./NavbarUser";
import {LogoThemed} from "../../routes/components/LogoThemed/LogoThemed";
import PropTypes from "prop-types";

export class DefaultNavbar extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <Navbar light expand="xs" fluid>
                <Nav navbar>
                    <NavItem className="navbar-brand d-lg-none">
                        <Link to="/">
                            <LogoThemed/>
                        </Link>
                    </NavItem>
                    <NavItem className="d-none d-md-block">
                        <span className="navbar-text">
                            <Link to="/">
                                <i className="fa fa-home"/>
                            </Link>
                        </span>
                        <span className="navbar-text px-2">
                            <i className="fa fa-angle-right"/>
                        </span>
                        <span className="navbar-text">
                            <Link to="/">Start</Link>
                        </span>
                        <span className="navbar-text px-2">
                            <i className="fa fa-angle-right"/>
                        </span>
                        <span className="navbar-text">
                            Page Link
                        </span>
                    </NavItem>
                </Nav>
                <Nav navbar className="ml-auto">
                    <NavbarUser logoutRequest={this.props.logoutRequest}/>
                </Nav>
            </Navbar>
        );
    }
}

DefaultNavbar.propTypes = {
    logoutRequest: PropTypes.func.isRequired,
};
