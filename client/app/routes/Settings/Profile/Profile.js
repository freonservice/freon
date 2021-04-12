import React from 'react';

import {
    Container,
    Row,
    Col,
    Input,
    Card,
    Button,
    CardBody,
    CardFooter,
    CardTitle,
    Label,
    FormGroup,
    Form
} from '../../../components';

import {HeaderMain} from '../../components/HeaderMain';
import * as PropTypes from 'prop-types';
import {connect} from 'react-redux';
import {
    profileRequest,
    updateUserProfileRequest
} from '../../../redux/profile/actions';

export class Profile extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            currentProfile: this.props.profile
        };
    }

    handleChange = (e) => {
        const {name, value} = e.target;
        const o = {...this.state.currentProfile, [name]: value};
        this.setState(function (previousState) {
            return {...previousState, currentProfile: o};
        });
    };

    handleUpdateProfile = (e) => {
        e.preventDefault();
        const {email, first_name, second_name} = this.state.currentProfile;
        if (email.trim() === '' || first_name.trim() === '' || second_name.trim() === '') {
            alert('Profile cant be empty');
            return;
        }
        this.props.updateUserProfileRequest(email, first_name, second_name);
    };

    render() {
        const {email, first_name, second_name} = this.state.currentProfile;
        return (
            <React.Fragment>
                <Container>
                    <HeaderMain
                        title="Profile Edit"
                        className="mb-5 mt-4"
                    />
                    <Row>
                        <Col lg={12}>
                            <Card>
                                <CardBody>
                                    <div className="d-flex mb-4">
                                        <CardTitle tag="h6">
                                            Profile Edit
                                        </CardTitle>
                                        <span className="ml-auto align-self-start small">
                                </span>
                                    </div>
                                    <Form>
                                        <FormGroup row>
                                            <Label for="firstName" sm={3} className="text-right">First Name</Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="first_name"
                                                    placeholder="First Name..."
                                                    onChange={(e) => this.handleChange(e)}
                                                    value={first_name}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="lastName" sm={3} className="text-right">Last Name</Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="second_name"
                                                    placeholder="Last Name..."
                                                    onChange={(e) => this.handleChange(e)}
                                                    value={second_name}
                                                />
                                            </Col>
                                        </FormGroup>
                                        <FormGroup row>
                                            <Label for="firstName" sm={3} className="text-right">Email</Label>
                                            <Col sm={8}>
                                                <Input
                                                    type="text"
                                                    name="email"
                                                    placeholder="Email..."
                                                    onChange={(e) => this.handleChange(e)}
                                                    value={email}
                                                />
                                            </Col>
                                        </FormGroup>
                                    </Form>
                                </CardBody>
                                <CardFooter className="text-right">
                                    <Button onClick={(e) => this.handleUpdateProfile(e)} color="primary">
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

Profile.propTypes = {
    profileRequest: PropTypes.func.isRequired,
    profile: PropTypes.shape({
        email: PropTypes.string,
        first_name: PropTypes.string,
        second_name: PropTypes.string,
    }).isRequired,
    errorMsg: PropTypes.string,
    updateUserProfileRequest: PropTypes.func.isRequired,
};

// Profile.defaultProps = {
//     profile: {}
// };

const mapStateToProps = (state) => ({
    profile: state.profile.profile,
});

const mapDispatchToProps = {
    profileRequest,
    updateUserProfileRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(Profile);