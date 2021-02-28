import { FC, useState, useEffect } from 'react'
import { Card, Form, Cascader, Input, Radio } from 'antd';
import { Category, getCategories} from 'models/Category'
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

  // FIXME 영균
  // useEffect(() => {
  //   getCategories().then((categories) => {
  //     console.log("is call once")
  //     setCategories(categories)
  //   })
  // }, [])

  const [form] = Form.useForm();
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
    console.log('Success:', values);
    setRefresh(true)
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
        onValuesChange={onRequiredTypeChange}
        onFinishFailed={onCategorySubmitFailed}
      >
        <Form.Item
          label="name"
          name="name"
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="code"
          name="code"
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="depth"
          name="depth"
        >
          <Radio.Group
            initialValues={depthDefaultValue}
            onChange={(e) => {
              setSelectedDepth(e.target.value)
            }}
          >
            {[
              { label: _const.CategoryDepth.Label.depth1, value: _const.CategoryDepth.Value.depth1},
              { label: _const.CategoryDepth.Label.depth2, value: _const.CategoryDepth.Value.depth2},
              { label: _const.CategoryDepth.Label.depth3, value: _const.CategoryDepth.Value.depth3},
              { label: _const.CategoryDepth.Label.depth4, value: _const.CategoryDepth.Value.depth4},
            ].map((v) => (
              <Radio.Button key={`depth-${v.value}`} value={v.value}>{v.label}</Radio.Button>
            ))}
          </Radio.Group>
        </Form.Item>
        <Form.Item
          label="parent_category"
          name="parent_category"
        >
          <Cascader
            options={parseCategoryToCascaderOptions(selectedDepth, categories)}
          />
        </Form.Item>
      </Form>
    </Card>
  )
}

const parseCategoryToCascaderOptions = (maximumDepth: string, categories: Array<Category>): Array<CascaderOptionType> => {
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
  if (maximumDepth === _const.CategoryDepth.Value.depth1) maxDepth = 1
  else if (maximumDepth === _const.CategoryDepth.Value.depth2) maxDepth = 2
  else if (maximumDepth === _const.CategoryDepth.Value.depth3) maxDepth = 3
  else if (maximumDepth === _const.CategoryDepth.Value.depth4) maxDepth = 4

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
  console.log(`depth: ${maximumDepth}, item.size: ${categories.length} => total cascade operating count: ${debug.count}`)
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