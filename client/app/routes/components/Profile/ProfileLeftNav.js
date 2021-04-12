import React from "react";
import {NavLink as RouterNavLink} from "react-router-dom";
import {
    Nav,
    NavItem,
    NavLink
} from "./../../../components";

const ProfileLeftNav = () => (
    <React.Fragment>
        <div className="mb-4">
            <Nav pills vertical>
                <NavItem>
                    <NavLink tag={RouterNavLink} to="/settings/profile-edit">
                        Profile Edit
                    </NavLink>
                </NavItem>
                <NavItem>
                    <NavLink tag={RouterNavLink} to="/settings/security-edit">
                        Security Edit
                    </NavLink>
                </NavItem>
            </Nav>
        </div>
    </React.Fragment>
);

export {ProfileLeftNav};
