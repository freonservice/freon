import React from 'react';

import {
    Container,
    Row,
    Card,
    Button,
    CardBody,
    Form, FormGroup, Label, Input, Col, CustomInput, CardFooter
} from '../../../components';

import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {HeaderMain} from '../../components/HeaderMain';
import {createUserRequest} from '../../../redux/users/actions';

export class CreateUser extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            user: {first_name: '', second_name: '', email: '', password: '', repeat_password: '', role: ''}
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        const o = {...this.state.user, [name]: value};
        this.setState({user: o});
    };

    handleCreateUser = () => {
        const {first_name, second_name, email, password, repeat_password, role} = this.state.user;
        this.props.createUserRequest(email, first_name, second_name, password, repeat_password, role);
        this.props.history.goBack();
    };

    render() {
        const {first_name, second_name, email, role} = this.state.user;

        return (
            <React.Fragment>
                <Container>
                    <HeaderMain
                        title="Create User"
                        className="mb-5 mt-4"
                    />
                    <Row>
                        <Col lg={12}>
                            <Card>
                                <CardBody>
                                    <Form>
                                        <div className="small mt-4 mb-3">
                                            Personal Information
                                        </div>
                                        <FormGroup row>
                                            <Label for="firstName" sm={3} className="text-right">
                                                First Name
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="first_name"
                                                    placeholder="First Name..."
                                                    value={first_name}
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="lastName" sm={3} className="text-right">
                                                Second Name
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="second_name"
                                                    placeholder="Last Name..."
                                                    value={second_name}
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="lastName" sm={3} className="text-right">
                                                <span className="text-danger">*</span> User Email
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="email"
                                                    placeholder="Email..."
                                                    value={email}
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <div className="small mt-5 mb-3">
                                            Availability
                                        </div>
                                        <FormGroup row>
                                            <Label for="jobType" sm={3} className="text-right">
                                                User Role
                                            </Label>
                                            <Col sm={8}>
                                                <CustomInput
                                                    type="select"
                                                    id="role"
                                                    name="role"
                                                    value={role}
                                                    onChange={(e) => this.handleChange(e)}
                                                >
                                                    <option value="">Select...</option>
                                                    <option value="admin">Admin</option>
                                                    <option value="translator">Translator</option>
                                                    <option value="moderator">Moderator</option>
                                                </CustomInput>
                                            </Col>
                                        </FormGroup>
                                        <div className="small mt-5 mb-3">
                                            Security
                                        </div>
                                        <FormGroup row>
                                            <Label for="newPassword" sm={3} className="text-right">
                                                <span className="text-danger">*</span> New Password
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="password"
                                                    name="password"
                                                    placeholder="Put your password here"
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row className="mb-0">
                                            <Label for="confirmNewPassword" sm={3} className="text-right">
                                                <span className="text-danger">*</span> Confirm New Password
                                            </Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="password"
                                                    name="repeat_password"
                                                    placeholder="Repeat you password"
                                                    onChange={(e) => this.handleChange(e)}
                                                />
                                            </Col>
                                        </FormGroup>
                                    </Form>
                                </CardBody>
                                <CardFooter className="text-right">
                                    <Button onClick={this.handleCreateUser} color="primary">
                                        Update Profile
                                    </Button>
                                </CardFooter>
                            </Card>
                        </Col>
                    </Row>
                </Container>
            </React.Fragment>
        );
    }
}

CreateUser.propTypes = {
    history: PropTypes.object.isRequired,
    createUserRequest: PropTypes.func.isRequired,
};

const mapDispatchToProps = {
    createUserRequest,
};

export default connect(null, mapDispatchToProps)(CreateUser);