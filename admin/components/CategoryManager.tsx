import React, { FC} from 'react'
import {
  Card,
  Form,
  Cascader,
  Input,
  InputNumber,
  notification,
  Button,
} from 'antd';
import { Category, getCategories, putCategories } from 'models/Category'
import { CascaderOptionType } from 'antd/lib/cascader';

type CategoryManagerProps = {
}

interface CategoryFormData {
  depth: string;
  sort: number;
  code: string;
  name: string;
  parent_category: string;
}

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 18 },
};

const CategoryManager: FC<CategoryManagerProps> = () => {
  const [categories, setCategories] = React.useState<Array<Category>>([]);

  React.useEffect(() => {
    getCategories().then((categories) => setCategories(categories))
    return function cleanup() {
      console.log("cleanup")
    };
  }, []);
  

  const onCategorySubmit = (c: CategoryFormData) => {
    if (!c) return

    const parentCategoryIDs = c.parent_category || []
    const parentID = parentCategoryIDs[parentCategoryIDs.length - 1]
    const parentCategory = categories.find((v) => v.ID === parentID)

    const category: Category = {
      ID: '',
      FullName: '',
      Status: '',
      Sort: c.sort,
      Name: c.name,
      Code: c.code,
      Depth: parentCategory?.Depth! + 1,
      Category1ID: parentCategory?.Category1ID,
      Category2ID: parentCategory?.Category2ID,
      Category3ID: parentCategory?.Category3ID,
      Category4ID: parentCategory?.Category4ID,
    }

    putCategories(category).then((res) => {
      openNotification(
        'success',
        '카테고리 등록 성공',
        (<>
          <p>Name: {res.Name}</p>
          <p>Code: {res.Code}</p>
          <p>Depth: {res.Depth}</p>
          <p>FullName: {res.FullName}</p>
        </>),
      )
      getCategories().then((categories) => {
        setCategories(categories)
      })
    }).catch((e) => {
      openNotification('error', "카테고리 관리등록 실패",  JSON.stringify(e))
    })
  };

  const onCategorySubmitFailed = (errorInfo: any) => {
    openNotification('error', "카테고리 관리등록 실패", JSON.stringify(errorInfo))
  };

    
  return (
    <Card title="카테고리 관리">
      <Form<CategoryFormData>
        {...layout}
        name="basic"
        initialValues={{ remember: true }}
        onFinish={onCategorySubmit}
        onFinishFailed={onCategorySubmitFailed}
      >
        <Form.Item
          label="parent_category"
          name="parent_category"
          rules={[{ type: 'array', required: true, message: 'parent_category 값을 선택해주세요' }]}
        >
          <Cascader
            options={parseCategoryToCascaderOptions(categories)}
            displayRender={label => label.join(' > ')}
            changeOnSelect
          />
        </Form.Item>
        <Form.Item
          label="name"
          name="name"
          rules={[{ type: 'string', required: true, message: 'name 값을 입력해주세요.' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="code"
          name="code"
          rules={[{ type: 'string', required: true, message: 'code 값을 입력해주세요.' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="sort"
          name="sort"
          initialValue={1}
          rules={[{ type: 'number', required: true, message: 'sort 값을 입력해주세요.' }]}
        >
          <InputNumber min={1} max={30} />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 6 }}>
          <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
            Submit
          </Button>
        </Form.Item>
      </Form>
    </Card>
  )
}

type NotiType = 'info' | 'success' | 'warn' | 'error' | 'open'
const openNotification = (notiType: NotiType, title: React.ReactNode, content: React.ReactNode) => {
  notification[notiType]({
    message: title,
    description: content,
    placement: "topRight",
  });
};


const parseCategoryToCascaderOptions = (categories: Array<Category>): Array<CascaderOptionType> => {
  const depth: Array<Map<string, Category>> = []
  categories.forEach((v) => {
    const d = depth[v.Depth - 1]
    if (d) {
      depth[v.Depth - 1].set(v.ID, v)
    } else {
      depth[v.Depth - 1] = new Map().set(v.ID, v)
    }
  })

  let maxDepth = 4

  const debug = {
    count: 0,
  }
  const result: Array<CascaderOptionType> = cascade({
    list: [],
    maxDepth,
    depth: depth,
    currentDepth: 1,
    parentId: '',
    debug: debug,
  })
  console.log(`item.size: ${categories.length} => total cascade operating count: ${debug.count}`)
  return result
}

type cascadeOption = {
  list: Array<CascaderOptionType>;
  maxDepth: number;
  depth: Array<Map<string, Category>>;
  currentDepth: number;
  parentId: string | undefined | number;
  debug?: any;
}

const cascade = ({ list, maxDepth, depth, currentDepth, parentId, debug }: cascadeOption): Array<CascaderOptionType> => {
  debug.count += 1
  if (currentDepth > maxDepth) return list
  if (!depth[currentDepth - 1]) return list

  depth[currentDepth - 1].forEach((cat) => {
    const one = list.find(({ value }) => value === cat.ID)
    if (one) {
      return cascade({
        list: one.children!,
        maxDepth,
        depth,
        currentDepth: currentDepth + 1,
        parentId: one.value,
        debug,
      })
    } else {
      let pid = ''
      if (cat.Depth === 2) pid = cat.Category1ID!
      else if (cat.Depth === 3) pid = cat.Category2ID!
      else if (cat.Depth === 4) pid = cat.Category3ID!

      if (pid === parentId) {
        list.push({
          label: cat.Name,
          value: cat.ID,
          children: [],
          disabled: cat.Depth === 4,
        })
        return cascade({
          list,
          maxDepth,
          depth,
          currentDepth,
          parentId,
          debug,
        })
      }
    }
  })
  return list
}

export default CategoryManager