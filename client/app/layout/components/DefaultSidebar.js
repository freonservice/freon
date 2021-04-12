import React from "react";
import {Link} from "react-router-dom";

import {Sidebar} from "./../../components";

import {UserSidebar} from "./UserSidebar";

import {LogoThemed} from "../../routes/components/LogoThemed/LogoThemed";
import SidebarProfile from "../../routes/components/Sidebar/SidebarProfile";

export const DefaultSidebar = () => (
    <Sidebar>
        <Sidebar.HideSlim>
            <Sidebar.Section>
                <Link to="/" className="sidebar__brand">
                    <LogoThemed checkBackground/>
                </Link>
            </Sidebar.Section>
        </Sidebar.HideSlim>

        <Sidebar.MobileFluid>
            <SidebarProfile/>
            <Sidebar.Section fluid cover>
                <UserSidebar/>
            </Sidebar.Section>
        </Sidebar.MobileFluid>
    </Sidebar>
);
