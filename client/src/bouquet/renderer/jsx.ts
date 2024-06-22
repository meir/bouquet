
export const JSX = {
  element: function (name: string & symbol, attrs: {[key: string]: string}, ...content: any) {
    const element: DocumentFragment = name === JSX.fragment ? 
      document.createDocumentFragment() :
      document.createElement(name)

    if (attrs) Object.keys(attrs).forEach(key => {
      if (key.startsWith('on')) {
        element[key] = attrs[key]
        return
      }
      
      if (element instanceof Element)
        element.setAttribute(key, attrs[key])
    })

    content.forEach((child: Element) => {
      if (Array.isArray(child)) {
        element.append(...child)
        return
      }
      element.append(child)
    })
    
    return element
  },
  fragment: Symbol('fragment')
}
