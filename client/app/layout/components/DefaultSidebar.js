import React from "react";

import {Sidebar} from "./../../components";

import {UserSidebar} from "./UserSidebar";

import SidebarProfile from "../../routes/components/Sidebar/SidebarProfile";

export const DefaultSidebar = () => (
    <Sidebar>
        <Sidebar.MobileFluid>
            <SidebarProfile/>
            <Sidebar.Section fluid cover>
                <UserSidebar/>
            </Sidebar.Section>
        </Sidebar.MobileFluid>
    </Sidebar>
);
