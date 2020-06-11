import React from "react";
import {Box, Form, Image, Text, FormField, Button, Anchor} from "grommet";
import {Link} from "react-router-dom";
import Select from 'react-select';
import axios from "axios";

class SignupView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            authorities: [],

            authority: null,
            signupAuthority: false,

            message: null,
            loading: false
        }

        this.signupAuthority = this.signupAuthority.bind(this);
        this.getAuthorities = this.getAuthorities.bind(this);
        this.selectAuthority = this.selectAuthority.bind(this);
        this.signup = this.signup.bind(this);
    }

    componentDidMount() {
        this.getAuthorities()
    }

    signupAuthority() {
        const { signupAuthority } = this.state;

        this.setState({
            signupAuthority: !signupAuthority,
            authority: null
        })
    }

    //
    // https://apiserver.edge-net.org/apis/apps.edgenet.io/v1alpha/authorities
    getAuthorities() {
// 'https://eapi-test.planet-lab.eu/apis/apps.edgenet.io/v1alpha/authorities'
        axios.get('https://eapi-test.planet-lab.eu/apis/apps.edgenet.io/v1alpha/authorities',
            {
                // withCredentials: true
                // headers: { Authorization: "Bearer " + anonymous_token }
            })
            .then(({data}) =>
                this.setState({
                    authorities: data.items.map(item => {
                        return {
                            value: item.metadata.name, label: item.spec.fullname + ' ('+item.spec.shortname+')'
                        }
                    }),
                }, () => console.log(data))
            );
    }

    selectAuthority(value) {
        console.log(value)
    }

    signup({value}) {
        console.log(value)
    }

    render() {
        const { logo, title } = this.props;
        const { authorities, authority, signupAuthority, message, loading } = this.state;

        return (
            <Box align="center">
                <Box gap="small" pad={{vertical:'large'}}>
                    {logo && <Image style={{maxWidth:'25%',margin:'50px auto'}} src={logo} alt={title} />}
                    {title ? title : "Signup"}
                </Box>
                <Form onSubmit={this.signup}>
                    <Box gap="medium" alignSelf="center" width="medium" alignContent="center" align="stretch">


                            <Box border={{side: 'bottom', color: 'brand', size: 'small'}}
                                 pad={{vertical: 'medium'}} gap="small"
                            >

                                {signupAuthority ?
                                    <Box>
                                        <Text color="dark-2">
                                            Please complete with the information of the institution you are part of
                                        </Text>
                                        <FormField label="Institution full name" name="fullname" required validate={{ regexp: /^[a-z]/i }} />
                                        <FormField label="Institution shortname or initials" name="shortname" required validate={{ regexp: /^[a-z]/i }} />
                                        <FormField label="Address" name="address" required validate={{ regexp: /^[a-z]/i }} />
                                        <FormField label="Web page" name="url" required validate={{ regexp: /^[a-z]/i }} />
                                    </Box>
                                    :
                                    <Select placeholder="Select your institution"
                                            isSearchable={true} isClearable={true}
                                            options={authorities}
                                        // value={}
                                            name=""
                                            onChange={this.selectAuthority}/>
                                }


                                <Anchor onClick={this.signupAuthority}>
                                    {signupAuthority ? "I want to select an existing institution" : "My institution is not on the list" }
                                </Anchor>
                            </Box>


                        <Box border={{side:'bottom',color:'brand',size:'small'}}>
                            <FormField label="Firstname" name="firstname" required validate={{ regexp: /^[a-z]/i }} />
                            <FormField label="Lastname" name="lastname" required validate={{ regexp: /^[a-z]/i }} />
                            <FormField label="E-Mail" name="email" required />
                            <FormField label="Phone" name="phone" required />

                            <Box direction="row" pad={{vertical:'medium'}} justify="between" align="center">
                                <Link to="/migrate">Migrate my PlanetLab Europe account</Link>
                                <Button type="submit" primary label="Signup" disabled={loading} />
                            </Box>
                        </Box>
                        <Box direction="row" >
                            <Link to="/">Go back</Link>
                        </Box>
                        {message && <Text color="status-critical">{message}</Text>}
                    </Box>
                </Form>
            </Box>
        )
    }
}

export default SignupView;