import { FC } from 'react'
import DefaultLayout from 'components/DefaultLayout'

type Props = {
  data?: any;
}


const IndexPage:FC<Props> = ({ data }) => {
  console.log(data)
  return (
    <DefaultLayout
      title="굳세어라 김치김치"
    >
      김치김치
    </DefaultLayout>
  )
}

// This gets called on every request
export async function getServerSideProps() {
  // Fetch data from external API
  const res = await fetch(`https://www.naver.com/data`)
  const data = await res.json()

  // Pass data to the page via props
  return { props: { data } }
}

export default IndexPage
