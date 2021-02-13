import { Fragment } from 'react'

import Head from 'next/head'

import styled, { createGlobalStyle } from 'styled-components'

const IndexPage = () => (
  <Fragment>
    <GlobalLayout />
    <Head>
      <title> hello world </title>
      <meta name="google-signin-scope" content="profile email"/>
      {/* <meta name="google-signin-client_id" content="YOUR_CLIENT_ID.apps.googleusercontent.com"/> */}
      <meta name="google-signin-client_id" content={`${process.env.GOOGLE_OAUTH_CLIENT_ID}.apps.googleusercontent.com`} />
      <script src="https://apis.google.com/js/platform.js" async defer></script>
    </Head>
    <CardLayout>
      <CardHeader>
        로그인
      </CardHeader>
      <CardBody>
        <div className="g-signin2" data-onsuccess="onSignIn" ></div>
      </CardBody>
    </CardLayout>
  </Fragment>
)


// const onSignIn = function(googleUser) {
//   // Useful data for your client-side scripts:
//   const profile = googleUser.getBasicProfile();
//   console.log("ID: " + profile.getId()); // Don't send this directly to your server!
//   console.log('Full Name: ' + profile.getName());
//   console.log('Given Name: ' + profile.getGivenName());
//   console.log('Family Name: ' + profile.getFamilyName());
//   console.log("Image URL: " + profile.getImageUrl());
//   console.log("Email: " + profile.getEmail());

//   // The ID token you need to pass to your backend:
//   var id_token = googleUser.getAuthResponse().id_token;
//   console.log("ID Token: " + id_token);
// }



const CardBody = styled.div`
  width: (100% - 1rem);
  height: calc(20rem - 3rem);
  padding: 0.5rem;
 
  border-radius: 0px 0px 6px 6px;
  background-color: red;
`

const CardHeader = styled.div`
  width: calc(100% - 1rem);
  display: flex;
  flex-direction: row;
  padding: 0 0.5rem;
  color: #333;
  font-size: 1.25rem;
  height: 3rem;
  line-height: 3rem;
  border-bottom: 1px solid #acacac;
`

const CardLayout = styled.div`
  width: 100%;
  color: #333;
  height: 100%;
  min-width: 20rem;
  min-height: 20rem;
  border: 1px solid #acacac;
  border-radius: 6px;
  box-shadow: 3px 3px 3px #ccc;
`


const GlobalLayout = createGlobalStyle`
  body {
    width: 100vw;
    height: 100vh;
    margin: 0;
    display: flex;
    align-items: center;
    justify-content: center;
  }
`




export default IndexPage
