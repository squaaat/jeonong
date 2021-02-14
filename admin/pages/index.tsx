import { FC } from 'react'
import DefaultLayout from 'components/DefaultLayout'

import { Session, MockSession } from 'models/Session'
import Card from 'components/Card'

import sg from 'utils/sample-data'

type PageProps = {
  session: Session;
}

type ServerProps = {
  props: PageProps;
}

const IndexPage:FC<PageProps> = ({ session }) => {
  console.log(session)
  return (
    <DefaultLayout
      session={session}
      title="굳세어라 김치김치"
    >
      <Card
        header="상품등록"
      >
        <p>
          김치김치
        </p>
      </Card>
    </DefaultLayout>
  )
}


// This function gets called at build time
export async function getStaticProps() {
  
  const data: ServerProps = {
    props: {
      session: MockSession,
    }
  }
  return data
}

export default IndexPage
