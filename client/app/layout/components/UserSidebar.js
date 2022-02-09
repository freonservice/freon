import React from 'react';

import {SidebarMenu} from './../../components';

export const UserSidebar = () => (
    <SidebarMenu>
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-home"/>}
            title="Dashboard"
            to='/dashboard'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-globe"/>}
            title="Localizations"
            to='/localizations'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-newspaper-o"/>}
            title="Categories"
            to='/categories'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-key"/>}
            title="Identifiers"
            to='/identifiers'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-bolt"/>}
            title="Translations"
            to='/translations'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-folder-open"/>}
            title="Translation Files"
            to='/translation-files'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-users"/>}
            title="Users List"
            to='/users'
        />
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-gear"/>}
            title="Settings"
        >
            <SidebarMenu.Item title="Profile" to='/settings/profile-edit' exact/>
            <SidebarMenu.Item title="Security" to='/settings/security-edit' exact/>
            <SidebarMenu.Item title="Storage" to='/settings/storage-edit' exact/>
            <SidebarMenu.Item title="Translation" to='/settings/translation-edit' exact/>
        </SidebarMenu.Item>

    </SidebarMenu>
);
