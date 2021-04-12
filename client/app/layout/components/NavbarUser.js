import React from 'react';
import PropTypes from 'prop-types';
import {NavItem} from './../../components';
import {Button} from "reactstrap";

export class NavbarUser extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <NavItem className="ml-2">
                <Button onClick={this.props.logoutRequest}>
                    <i className="fa fa-power-off"/>
                </Button>
            </NavItem>
        );
    }
}

NavbarUser.propTypes = {
    logoutRequest: PropTypes.func.isRequired,
};
