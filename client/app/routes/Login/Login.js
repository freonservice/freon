import React from "react";
import {
    Form,
    FormGroup,
    FormText,
    Input,
    CustomInput,
    Button,
    Label,
    EmptyLayout,
    ThemeConsumer
} from "../../components";

import {HeaderAuth} from "../components/Pages/HeaderAuth";
import {FooterAuth} from "../components/Pages/FooterAuth";
import * as PropTypes from "prop-types";
import {loginRequest} from "../../redux/login/actions";
import {connect} from "react-redux";
import {DangerAlert} from "../components/Alert/danger";

class Login extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            email: "",
            password: "",
            isLoading: false,
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        this.setState({
            ...this.state, [name]: value
        });
    };

    handleSubmit = (e) => {
        e.preventDefault();
        const {email, password} = this.state;
        this.props.loginRequest(email.trim(), password.trim());
        this.setState({email: "", password: ""});
    };

    render() {
        return (
            <EmptyLayout>
                <EmptyLayout.Section center>
                    <HeaderAuth
                        title="Sign In to Application"
                    />
                    <Form className="mb-3">
                        <FormGroup>
                            <Label for="emailAddress">
                                Email Address
                            </Label>
                            <Input
                                type="email"
                                name="email"
                                id="emailAddress"
                                placeholder="Enter email..."
                                className="bg-white"
                                value={this.state.email}
                                onChange={this.handleChange}
                            />
                            <FormText color="muted">
                                We&amp;ll never share your email with anyone else.
                            </FormText>
                        </FormGroup>
                        <FormGroup>
                            <Label for="password">
                                Password
                            </Label>
                            <Input
                                type="password"
                                name="password"
                                id="password"
                                placeholder="Password..."
                                className="bg-white"
                                value={this.state.password}
                                onChange={this.handleChange}
                            />
                        </FormGroup>
                        <FormGroup>
                            <CustomInput
                                type="checkbox"
                                id="rememberPassword"
                                label="Remember Password"
                                inline
                            />
                        </FormGroup>
                        {this.props.loginState.error && <DangerAlert text={this.props.loginState.error}/>}
                        <ThemeConsumer>
                            {
                                ({color}) => (
                                    <Button onClick={this.handleSubmit} color={color} block>
                                        Sign In
                                    </Button>
                                )
                            }
                        </ThemeConsumer>
                    </Form>
                    <FooterAuth/>
                </EmptyLayout.Section>
            </EmptyLayout>
        );
    }
}

Login.propTypes = {
    loginRequest: PropTypes.func.isRequired,
    loginState: PropTypes.object.isRequired,
};

const mapStateToProps = state => ({
    loginState: state.login,
});

export default connect(mapStateToProps, {loginRequest})(Login);
