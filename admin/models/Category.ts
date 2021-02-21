import { getCategories, CategoryFromServer } from 'api/rest'

import { Tree } from 'utils/tree'

export type Category = {
  id: string;
  status: string;
  name: string
  
  depth1?: CategoryDepth;
  depth2?: CategoryDepth;
  depth3?: CategoryDepth;
}

type CategoryDepth = {
  id: string;
  name: string;
  status: string;
  keyword: Keyword;
  parentKeyword: Keyword;
}

type Keyword = {
  id: string;
  name: string;
  code: string;
}

export const GetCategories = async (): Promise<Category[]> => {
  const result = await getCategories()
  if (result.success === false) {
    throw result.error
  }
  const dataAry = result.GetCategoriesResult || []

  const categoryDepths: CategoryDepth[] = []
  dataAry.forEach((v: CategoryFromServer) => {
    categoryDepths.push({
      id: v.ID,
      status: v.Status,
      name: v.Keyword.Name,
      keyword: {
        id: v.Keyword.ID,
        name: v.Keyword.Name,
        code: v.Keyword.Code,
      },
      parentKeyword: {
        id: v.ParentKeyword.ID,
        name: v.ParentKeyword.Name,
        code: v.ParentKeyword.Code,
      },
    })
  })

  // FIXME: 키워드로 연결하려니까
  // '가방' 카테고리 depth3은 depth2에 대해서 남성의류, 여성의류
  // 두 곳의 부모로 키워드가 있는데, 카테고리와 키워드가 동일선상에서 부모, 자식관계를 찾는바람에
  // 데이터 로직이 전부 꼬여버림.
  // 카테고리 설계를 다시 할 필요가 있음.
  console.log('---------1---------')
  const t = new Tree({
    list: categoryDepths,
    idPath: 'id',
    parentIdPath: 'parent.id',
  }).build()
  

  console.log(t.toString())
  console.log('---------2---------')

  // for(const catID in mapSourceCategories) {
  //   const cd: CategoryDepth = mapSourceCategories[catID] || {}

  //   // 최상위 카테고리 depth1
  //   if (cd.keyword.id === cd.parentKeyword.id) {
  //     tree[cd.id] = cd
  //   }
  // }

  return []
}

// input: 
// Tree(list: Categories, self.id: 'KeywordID', self.parentId: 'ParentKeywordID')
// output: Tree
//   Tree.toString(depth: 0) // printout all
//   Tree.toString(depth: 3) // depth 3까지 쭉 출력하도록