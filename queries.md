## Query - 1         
query{
	getCPTemplate(id: 1) {
    id
    definition
    name
  }
}

## Query - 2
query{
	getCPTemplates(limit: 2) {
    id
    name
    definition
  }
}

## Query - 3
query {
  getCPWorkbook (id: 1) {
    id
    scope
    status
    last_modified_by {
      userId
    }
    template {
      id
      name
    }
    comments {
      id
      comment
    }
  }
}

## Query - 4
query {
  getCPWorkbooks (limit: 3) {
    id
    scope
    status
    last_modified_by {
      userId
    }
    template {
      id
      name
    }
    comments {
      id
      comment
    }
  }
}

## Query - 5
query {
	getCPPin(id : 3) {
  	id
    title
    description
    workbook {
      id
      status
      scope
    }
}
}

## Query - 6
query {
	getCPPins(limit : 3) {
  	id
    title
    description
    workbook {
      id
      status
      scope
    }
}
}

## Query - 7
query {
	getCPWorkbookComments(workbookID: 1) {
    id
    workbookID
    comment
  }
}

## Query - 8
query {
	getCPWorkbookComments(workbookID: 1, limit: 2) {
    id
    workbookID
    comment
  }
}

## Query - 9
mutation {
  updateCPWorkbook ( 
    workbookID: 1,
    Status : PUBLISHED,
  ) {
    id
    status
  }
}

## mutation - 1
mutation {
addCPWorkbookComment (
  workbookID : 1
  comment: "Comment-4"
) 
}

## mutation - 2
mutation {
  addCPPin(
    pin: {
      title: "second_pin"
      description: "second_pin_desc"
      filters: "{\"name\": \"Praveen\", \"age\": \"40\", \"city\": \"Denver\"}"
      context: "{\"name\": \"Sadique\", \"age\": \"30\", \"city\": \"Lucknow\"}"
      workbookID: "2"
      visualization_flag: true
    }
  ) {
    id
    title
  }
}

## mutation - 3
mutation {
  updateCPWorkbook(workbookID: 5, Status: PUBLISHED) {
    id
  }
}
