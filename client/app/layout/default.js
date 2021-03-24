import React from "react";
import PropTypes from "prop-types";

import {
    Layout,
    ThemeProvider,
} from "./../components";

import "./../styles/bootstrap.scss";
import "./../styles/main.scss";
import "./../styles/plugins/plugins.scss";
import "./../styles/plugins/plugins.css";

import {DefaultSidebar} from "./components/DefaultSidebar";
import {DefaultNavbar} from "./components/DefaultNavbar";
import {connect} from "react-redux";
import {logoutRequest} from "../redux/login/actions";

const favIcons = [
    {rel: "icon", type: "image/x-icon", href: require("./../images/favicons/favicon.ico")},
    {rel: "apple-touch-icon", sizes: "180x180", href: require("./../images/favicons/apple-touch-icon.png")},
    {rel: "icon", type: "image/png", sizes: "32x32", href: require("./../images/favicons/favicon-32x32.png")},
    {rel: "icon", type: "image/png", sizes: "16x16", href: require("./../images/favicons/favicon-16x16.png")}
];

class AppLayout extends React.Component {
    render() {
        const {children, logoutRequest} = this.props;

        return (
            <ThemeProvider initialStyle="light" initialColor="primary">
                <Layout sidebarSlim favIcons={favIcons}>
                    <Layout.Navbar>
                        <DefaultNavbar logoutRequest={logoutRequest}/>
                    </Layout.Navbar>
                    <Layout.Sidebar>
                        <DefaultSidebar/>
                    </Layout.Sidebar>
                    <Layout.Content>
                        {children}
                    </Layout.Content>
                </Layout>
            </ThemeProvider>
        );
    }
}

AppLayout.propTypes = {
    children: PropTypes.node.isRequired,
    logoutRequest: PropTypes.func.isRequired,
};

export default connect(null, {logoutRequest})(AppLayout);
