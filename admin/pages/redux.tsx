import React from 'react';
import { useSelector } from 'react-redux';
import {NextPage} from 'next';
import {wrapper, State} from 'store/index';

export const getServerSideProps = wrapper.getServerSideProps(
    ({store, req, res, ...etc}) => {
        console.log('2. Page.getServerSideProps uses the store to dispatch things');
        store.dispatch({type: 'TICK', payload: 'was set in other page'});
    }
);


// export const getStaticProps = wrapper.getStaticProps(
//     ({store, preview}) => {
//         console.log('2. Page.getStaticProps uses the store to dispatch things');
//         store.dispatch({type: 'TICK', payload: 'was set in other page ' + preview});
//     }
// );


// Page itself is not connected to Redux Store, it has to render Provider to allow child components to connect to Redux Store
const Page: NextPage = () => {
  const { tick } = useSelector<State, State>(state => state)
  console.log(tick)
  return (
    <div>{tick.tick}</div>
  );
}
// you can also use Redux `useSelector` and other hooks instead of `connect()`
export default Page