import {
    Button,
    Card,
    CardBody,
    CardTitle,
    Col,
    Form,
    FormGroup,
    Input,
    CustomInput,
    Row
} from '../../components';
import React from 'react';
import PropTypes from 'prop-types';
import {Typeahead} from 'react-bootstrap-typeahead';
import {toast} from 'react-toastify';
import {Media} from 'reactstrap';
import TagsInput from '../../components/TagsInput';

const defaultPlatforms = {
    'web': {id: 0, label: 'web', name: 'web', checked: true},
    'apple': {id: 1, label: 'apple', name: 'apple', checked: false},
    'android': {id: 2, label: 'android', name: 'android', checked: false},
};

export class Action extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            chosenCategory: [], // selected category using like array
            platforms: defaultPlatforms,
        };
    }

    handleChangeCheckboxes = (e) => {
        const {checked, name} = e.target;
        const platforms = {...this.state.platforms};
        platforms[name]['checked'] = checked;
        if (checked) {
            this.props.chosenIdentifier.platforms.push(name);
        } else {
            this.props.chosenIdentifier.platforms = this.props.chosenIdentifier.platforms.filter(x => x !== name);
        }
        this.setState(function (previousState) {
            return {...previousState, platforms: platforms};
        });
    };

    handleChangeCategory = (category) => {
        this.setState({...this.state, chosenCategory: category});
    };

    handleResetChosenIdentifier = () => {
        this.props.handleResetChosenIdentifier();
        this.setState({
            chosenCategory: [],
            platforms: {...defaultPlatforms},
            namedList: []
        });
    };

    handleSubmitIdentifier = (e) => {
        e.preventDefault();
        const {id, name, description, example_text} = this.props.chosenIdentifier;
        const {namedList} = this.state;

        let categoryId = 0;
        if (this.state.chosenCategory.length > 0) {
            categoryId = this.state.chosenCategory[0].id;
        }
        if (name.trim() === '') {
            toast.success(contentSuccess('Value not valid!'));
            return;
        }

        const newPlatforms = [];
        if (this.state.platforms['web'].checked) {
            newPlatforms.push('web');
        }
        if (this.state.platforms['apple'].checked) {
            newPlatforms.push('apple');
        }
        if (this.state.platforms['android'].checked) {
            newPlatforms.push('android');
        }

        if (id > 0) {
            this.props.updateIdentifierRequest(id, name, description, example_text, categoryId, newPlatforms, namedList);
        } else {
            this.props.createIdentifierRequest(name, description, example_text, categoryId, newPlatforms, namedList);
        }
        this.handleResetChosenIdentifier();
    };

    render() {
        const {
            chosenIdentifier,
            listCategories,
            handleChangeIdentifierInformation
        } = this.props;
        const isEdit = chosenIdentifier.id > 0;
        const {platforms, chosenCategory, namedList} = this.state;

        if (isEdit) {
            platforms['web'].checked = false;
            platforms['apple'].checked = false;
            platforms['android'].checked = false;
            chosenIdentifier.platforms.forEach(function (platform) {
                // eslint-disable-next-line no-prototype-builtins
                if (platforms.hasOwnProperty(platform)) {
                    platforms[platform].checked = true;
                }
            });
        }

        return (
            <Card className="mb-3">
                <CardBody>
                    <CardTitle tag="h6" className="mb-4">
                        {isEdit ? (
                            <Row>
                                <Col sm={12} lg={7}>
                                    Edit Identifier
                                </Col>
                                <Col sm={12} lg={3}>
                                    <Button color="primary" onClick={this.handleResetChosenIdentifier}>Reset
                                        Form</Button>
                                </Col>
                            </Row>
                        ) : (`Add Identifier`)}
                    </CardTitle>
                    <Form>
                        <FormGroup>
                            <Input
                                type="text"
                                value={chosenIdentifier.name}
                                onChange={(e) => handleChangeIdentifierInformation(e)}
                                name="name"
                                placeholder="Key (ex: HELLO_WORLD_MESSAGE)"
                            />
                        </FormGroup>
                        <FormGroup>
                            <Typeahead
                                clearButton
                                id="list-category-id"
                                defaultSelected={listCategories.slice(0, 5)}
                                labelKey="name"
                                options={listCategories}
                                placeholder="Choose a category..."
                                onChange={(selected) => this.handleChangeCategory(selected)}
                                selected={chosenCategory}
                            />
                        </FormGroup>
                        <FormGroup>
                            <Input
                                type="textarea"
                                value={chosenIdentifier.description}
                                onChange={(e) => handleChangeIdentifierInformation(e)}
                                name="description"
                                placeholder="Short description (Optional)"
                            />
                        </FormGroup>
                        <FormGroup>
                            <Input
                                type="textarea"
                                value={chosenIdentifier.example_text}
                                onChange={(e) => handleChangeIdentifierInformation(e)}
                                name="example_text"
                                placeholder="Example translated text (Optional)"
                            />
                        </FormGroup>
                        <FormGroup className="mt-2">
                            {
                                Object.entries(platforms)
                                    .map(([key, value]) =>
                                        <CustomInput
                                            {...value}
                                            key={value.name}
                                            onChange={this.handleChangeCheckboxes}
                                            type="checkbox"
                                            id={`basic-behaviors-${key}`}
                                        />
                                    )
                            }
                        </FormGroup>
                        <Button onClick={(e) => this.handleSubmitIdentifier(e)} color="primary">
                            {
                                isEdit ? (`Edit`) : (`Create`)
                            }
                        </Button>
                    </Form>
                </CardBody>
            </Card>
        );
    }
}

Action.propTypes = {
    listCategories: PropTypes.array.isRequired,
    handleChangeIdentifierInformation: PropTypes.func.isRequired,
    chosenIdentifier: PropTypes.object.isRequired,
    createIdentifierRequest: PropTypes.func.isRequired,
    updateIdentifierRequest: PropTypes.func.isRequired,
    handleResetChosenIdentifier: PropTypes.func.isRequired,
};

Action.defaultProps = {
    listCategories: [],
};

// eslint-disable-next-line react/prop-types,no-unused-vars
const contentSuccess = (description) => (
    <Media>
        <Media middle left className="mr-3">
            <i className="fa fa-fw fa-2x fa-check"/>
        </Media>
        <Media body>
            <Media heading tag="h6">
                Success!
            </Media>
            <p>{description}</p>
        </Media>
    </Media>
);
