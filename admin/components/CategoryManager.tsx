import { FC, useState, useEffect } from 'react'
import { Card, Form, Cascader, Input, InputNumber, Radio, Button } from 'antd';
import { Category, getCategories, putCategories } from 'models/Category'
import { CascaderOptionType } from 'antd/lib/cascader';

import _const from 'utils/const'

type CategoryManagerProps = {
}

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 18 },
};

type RequiredMark = boolean | 'optional';
const depthDefaultValue = _const.CategoryDepth.Value.depth2

const CategoryManager: FC<CategoryManagerProps> = () => {
  const [categories, setCategories] = useState(Array<Category>());
  const [isRefresh, setRefresh] = useState(true);
  const [selectedDepth, setSelectedDepth] = useState(depthDefaultValue)

  // const [form] = Form.useForm();
  const [requiredMark, setRequiredMarkType] = useState<RequiredMark>('optional');
  const onRequiredTypeChange = ({ requiredMarkValue }: { requiredMarkValue: RequiredMark }) => {
    setRequiredMarkType(requiredMarkValue);
  };
  // console.log(form)
  // console.log(requiredMark)

  useEffect(() => {
    if(isRefresh) {
      getCategories().then((categories) => {
        console.log("is call once")
        setCategories(categories)
        setRefresh(false)
      })
    }
    return function cleanup() {
      console.log("cleanup")
    };
  }, [isRefresh]);


  const onCategorySubmit = (values: any) => {
    if (!values) return
    let depth = 1
    if (values?.depth === _const.CategoryDepth.Label.depth1) {
      depth = 1
    } else if (values?.depth === _const.CategoryDepth.Label.depth2) {
      depth = 2
    } else if (values?.depth === _const.CategoryDepth.Label.depth3) {
      depth = 3
    }

    const c: Category = {
      ID: '',
      FullName: '',
      Status: '',
      Sort: values?.sort,
      Name: values?.name,
      Code: values?.code,
      Depth: depth,
    }
    const cats = values?.parent_category || []

    const parentID = cats[cats.length - 1]
    const pc = categories.find((v) => v.ID === parentID)

    console.log(values)
    console.log(values?.parent_category)
    console.log(pc)
    if (depth > 1) {
      const pc = categories.find((v) => v.ID === values?.parent_category)
      c.Category1ID = pc?.Category1ID
      c.Category2ID = pc?.Category2ID
      c.Category3ID = pc?.Category3ID
      c.Category4ID = pc?.Category4ID
    }
  
    console.log(c)
    putCategories(c).then((res) => {
      console.log(res)
      setRefresh(true)
    }).catch((e) => {
      console.log(e)
    })
  };

  const onCategorySubmitFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <Card title="카테고리 관리">
      <Form
        {...layout}
        name="basic"
        initialValues={{ remember: true }}
        onFinish={onCategorySubmit}
        onFinishFailed={onCategorySubmitFailed}
        onValuesChange={onRequiredTypeChange}
        >
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
          rules={[{ type: 'number', required: true, message: 'sort 값을 입력해주세요.' }]}
        >
          <InputNumber min={1} max={30} defaultValue={1} />
        </Form.Item>
        <Form.Item
          label="depth"
          name="depth"
          rules={[{ type: 'string', required: true, message: 'depth 값을 입력해주세요.' }]}
        >
          <Radio.Group
            defaultValue={depthDefaultValue}
            onChange={(e) => {
              setSelectedDepth(e.target.value)
            }}
          >
            {[
              { label: _const.CategoryDepth.Label.depth1, value: _const.CategoryDepth.Value.depth1},
              { label: _const.CategoryDepth.Label.depth2, value: _const.CategoryDepth.Value.depth2},
              { label: _const.CategoryDepth.Label.depth3, value: _const.CategoryDepth.Value.depth3},
            ].map((v) => (
              <Radio.Button key={`depth-${v.value}`} value={v.value}>{v.label}</Radio.Button>
            ))}
          </Radio.Group>
        </Form.Item>
        <Form.Item
          label="parent_category"
          name="parent_category"
          rules={selectedDepth === _const.CategoryDepth.Value.depth1 ? [] : [{ type: 'array', required: true, message: 'parent_category 값을 선택해주세요' }]}
        >
          <Cascader
            options={parseCategoryToCascaderOptions(selectedDepth, categories)}
          />
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

const parseCategoryToCascaderOptions = (selectedDepth: string, categories: Array<Category>): Array<CascaderOptionType> => {
  const depth: Array<Map<string, Category>> = []
  categories.forEach((v) => {
    const d = depth[v.Depth - 1]
    if (d) {
      depth[v.Depth - 1].set(v.ID, v)
    } else {
      depth[v.Depth - 1] = new Map().set(v.ID, v)
    }
  })

  let maxDepth = 1
  if (selectedDepth === _const.CategoryDepth.Value.depth1) return []
  else if (selectedDepth === _const.CategoryDepth.Value.depth2) maxDepth = 1
  else if (selectedDepth === _const.CategoryDepth.Value.depth3) maxDepth = 2
  else if (selectedDepth === _const.CategoryDepth.Value.depth4) maxDepth = 3

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
  console.log(`depth: ${selectedDepth}, item.size: ${categories.length} => total cascade operating count: ${debug.count}`)
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
  if (currentDepth > maxDepth) {
    return list
  }
  if (!depth[currentDepth - 1]) {
    return list
  }
  depth[currentDepth - 1].forEach((cat) => {
    const one = list.find((item) => {
      return item.value === cat.ID
    })
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
      if (cat.Depth === 2) {
        pid = cat.Category1ID!
      } else if (cat.Depth === 3) {
        pid = cat.Category2ID!
      } else if (cat.Depth === 4) {
        pid = cat.Category3ID!
      }
      if (pid === parentId) {
        list.push({
          label: cat.Name,
          value: cat.ID,
          children: []
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