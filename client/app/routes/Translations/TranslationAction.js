import React from 'react';

import {
    Container,
    Row,
    Col,
    Card,
    Button,
    CardBody,
    Form, FormGroup, Label, Input
} from './../../components';

import PropTypes from 'prop-types';
import {HeaderMain} from '../components/HeaderMain';
import {withRouter} from 'react-router';
import {connect} from 'react-redux';
import {updateTranslationRequest} from '../../redux/translations/actions';

export class TranslationAction extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            translation: this.props.location.state,
            selectionStart: 0,
        };
    }

    handleChange = (e) => {
        e.preventDefault();
        const {name, value} = e.target;
        const o = {...this.state.translation, [name]: value};
        this.setState({...this.state, translation: o});
    };

    handleSaveButton = () => {
        const {id, singular, plural} = this.state.translation;
        console.log(id, singular, plural)
        this.props.updateTranslationRequest(id, singular.trim(), plural.trim());
        this.props.history.goBack();
    };

    render() {
        const {translation} = this.state;
        return (
            <React.Fragment>
                <Container>
                    <HeaderMain
                        title="Edit Translation"
                        className="mb-5 mt-4"
                    />
                    <Row>
                        <Col lg={8}>
                            <Card className="mb-3">
                                <CardBody>
                                    <div className="d-flex mb-5">
                                        <Button outline color="secondary" onClick={() => {
                                            this.props.history.goBack();
                                        }}>
                                            <i className="fa fa-fw fa-arrow-left"/>Back
                                        </Button>
                                        <Button color="success" className="ml-auto" onClick={this.handleSaveButton}>
                                            Save
                                        </Button>
                                    </div>
                                    <Form>
                                        <FormGroup>
                                            <Label for="text">
                                                Singular
                                            </Label>
                                            <Input
                                                type="textarea"
                                                name="singular"
                                                value={translation.singular}
                                                placeholder="Enter Your Message..."
                                                className="mb-2"
                                                onChange={(e) => this.handleChange(e)}
                                            />
                                        </FormGroup>
                                        <FormGroup>
                                            <Label for="text">
                                                Plural
                                            </Label>
                                            <Input
                                                type="textarea"
                                                name="plural"
                                                value={translation.plural}
                                                placeholder="Enter Your Message..."
                                                className="mb-2"
                                                onChange={(e) => this.handleChange(e)}
                                            />
                                        </FormGroup>
                                    </Form>
                                </CardBody>
                            </Card>
                        </Col>
                        <Col lg={4}>
                            <Card>
                                <CardBody>
                                    <h5><strong>Identifier info</strong></h5>
                                    <div className="mt-4 mb-2">
                                        <span className="small">
                                            <strong>Language code</strong>
                                        </span>
                                    </div>
                                    <p className="text-left">{translation.localization.locale}</p>

                                    <div className="mt-4 mb-2">
                                        <span className="small">
                                            <strong>Localization Key</strong>
                                        </span>
                                    </div>
                                    <p className="text-left">{translation.identifier.name}</p>

                                    <div className="mt-4 mb-2">
                                        <span className="small">
                                            <strong>Description</strong>
                                        </span>
                                    </div>
                                    <p className="text-left">{translation.description}</p>

                                    <div className="mt-4 mb-2">
                                        <span className="small">
                                            <strong>Example translation</strong>
                                        </span>
                                    </div>
                                    <p className="text-left">{translation.example_text}</p>
                                </CardBody>
                            </Card>
                        </Col>
                    </Row>
                </Container>
            </React.Fragment>
        );
    }
}

TranslationAction.propTypes = {
    history: PropTypes.object.isRequired,
    location: PropTypes.object.isRequired,
    updateTranslationRequest: PropTypes.func.isRequired,
};

const mapDispatchToProps = {
    updateTranslationRequest,
};

export default connect(null, mapDispatchToProps)(withRouter(TranslationAction));