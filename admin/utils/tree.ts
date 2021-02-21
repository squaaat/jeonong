type TreeOptions = {
  list: Array<object>;
  idPath: string
  parentIdPath: string
}

const defaultRootID: string = 'tree_default_root_id'
const defaultRootNode: any = {}

export class Tree {
  private list: Array<Node>
  private idPath: string
  private parentIdPath: string
  private rootNode?: Node = undefined

  constructor(opt: TreeOptions) {
    this.idPath = opt.idPath
    this.parentIdPath = opt.parentIdPath


    const mapNodes: Map<string, Node> = new Map()
    opt.list.forEach((item) => {
      const id: string = deepGet(item, this.idPath)
      const parentId = deepGet(item, this.parentIdPath)
      const node = new Node(id, item)
      node.parentId = parentId
      mapNodes.set(id, node)
    })

    const list: Array<Node> = []
    mapNodes.forEach((v) => {
      v.parent = mapNodes.get(v.parentId || '')
      list.push(v)
    })
    this.list = list
  }

  build(): Tree {
    this.rootNode = this.getRoot()
    return this
  }

  // root노드 선정기준
  // 1. id == parentId
  // 2. parentId가 없으면 rootNode
  // 3. parentId가 있지만, 그에 대한 Node가 없으면 해당 parentId를 가지고 RootNode를 만듬
  // 4. 1,2중에 해당하는 rootNode가 여러개라면 가공의 RootNode를 추가함
  getRoot(): Node {
    let rootCandidates: Array<Node> = []
    this.list.forEach((v) => {
      // 1번 케이스
      if (v.id === v.parentId) {
        rootCandidates.push(v)
      // 2번 케이스
      } else if (!v.parentId) {
        rootCandidates.push(v)
      // 3번 케이스
      } else if (v.parentId && !v.parent) {
        rootCandidates = [new Node(v.parentId, defaultRootNode)]
      }
    })
    if (rootCandidates.length === 1) {
      return rootCandidates[0]
    } else if (rootCandidates.length > 2) {
      const rootNode = new Node(defaultRootID, defaultRootNode)
      this.list = this.list.map((v) => {
        const item = rootCandidates.find((candidtate) => candidtate.id === v.id)
        if (item) {
          item.parentId = rootNode.id
          item.parent = rootNode
          return item
        }
        return v
      })
      this.list.push(rootNode)
      return rootNode
    }
    return new Node(defaultRootID, defaultRootNode)
  }

  toString(): string {
    if (this.rootNode === undefined) {
      throw Error('You must build() first')
    }
    const mapNodes: Map<string, Node> = new Map()
    this.list.forEach((item) => {
      mapNodes.set(item.id,  item)
    })

    const mapParentNodes: Map<string, Array<Node>> = new Map()
    this.list.forEach((item) => {
      const parentId = item.parentId || ''
      const node = mapNodes.get(parentId)
      const parentNodes = mapParentNodes.get(parentId) || []
      parentNodes.push(item)

      mapParentNodes.set(parentId, parentNodes)
    })

    const rootId = this.rootNode.id || ''
    const depth1 = mapParentNodes.get(rootId) || {}
    

    console.log(this.list.length)
    console.log(rootId)
    console.log(depth1.length)
    return ''
  }
}

class Node {
  id: string
  self: any

  parentId?: string = undefined
  parent?: Node = undefined

  constructor(id: string, item: any) {
    this.id = id
    this.self = item
  }
  setParent(id: string, n: Node): Node {
    this.parentId = id
    this.parent = n
    return this
  }

  value(): any {
    return this.self
  }

  hasNext(): boolean {
    return !!this.parent
  }

  next(): Node {
    return this.parent ? this.parent : this.self
  }
}

// Guard
function isString<T = any>(str: string | T): str is string {
  return typeof str === "string";
}

const deepGet = (
  obj: any,
  keys: string | (string | number)[],
  delimiter = "."
) => {
  if (isString(keys)) {
    keys = keys.split(delimiter);
  }
  return keys.reduce((xs, x) => (xs && xs[x] ? xs[x] : null), obj);
};