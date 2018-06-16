


function handleSearchResult(data) {

    if (!data) {
        return;
    }
     
    //console.info(data);
    var nodes = $('#tree').treeview('getEnabled');


    for (j in data) {

        checkSearchResult(data[j].Id);
    }
}


function checkSearchResult(id){
    if (!id){
        return;
    }

    var nodes = $('#tree').treeview('getEnabled');
    //console.info(nodes);
    for (i in nodes) {
        if (id == nodes[i].id) {
           
            $('#tree').treeview('selectNode', [ nodes[i], { silent: true } ]);
            break;
        }
    }

    return false;
}

function initTree() {

    $.post('/system/menu/list',
    function(_data) {

        $('#tree').treeview({

            data: generateTree(_data),
            highlightSearchResults:true,
            //multiSelect:true,
            nodeIcon: "glyphicon glyphicon-unchecked",
            collapseIcon: "glyphicon glyphicon-triangle-bottom",
            expandIcon: 'glyphicon glyphicon-triangle-right',
            backColor: '#F9F9F9',
            borderColor: '#C6C6C6',
            selectedIcon: 'glyphicon glyphicon-check',
            levels: 5,

            /* onNodeSelected:function(event,data){

                  checkParentNode(data.pid,data.id);
                }*/
            //showCheckbox:true,
        });
    },
    'json')

}

function initParentMenu() {

    $.post('/system/menu/list',
    function(_data) {

        $('#my-parentMenus').treeview({

            data: generateTree(_data),
            //highlightSearchResults:true,
            //multiSelect:true,
            nodeIcon: "glyphicon glyphicon-unchecked",
            collapseIcon: "glyphicon glyphicon-triangle-bottom",
            expandIcon: 'glyphicon glyphicon-triangle-right',
            backColor: '#F9F9F9',
            borderColor: '#C6C6C6',
            selectedIcon: 'glyphicon glyphicon-check',
            levels: 5,

            /* onNodeSelected:function(event,data){

                  checkParentNode(data.pid,data.id);
                }*/
            //showCheckbox:true,
        });
    },
    'json')

}

function initTreeMul() {

    $.post('/system/menu/list',
    function(_data) {

        $('#tree').treeview({

            data: generateTree(_data),
            highlightSearchResults:true,
            searchResultBackColor:'#D3AE30',
            multiSelect: true,
            nodeIcon: "glyphicon glyphicon-unchecked",
            collapseIcon: "glyphicon glyphicon-triangle-bottom",
            expandIcon: 'glyphicon glyphicon-triangle-right',
            backColor: '#F9F9F9',
            borderColor: '#C6C6C6',
            selectedIcon: 'glyphicon glyphicon-check',
            levels: 5,

            /* onNodeSelected:function(event,data){

                  checkParentNode(data.pid,data.id);
                }*/
            //showCheckbox:true,
        });
        //$('#tree').treeview('search', [0,{ ignoreCase: false, exactMatch: false }]);
    },
    'json')

}

function initRoleMenusTreeMul() {

    $.post('/system/menu/list',
    function(_data) {

        $('#roleMenus').treeview({

            data: generateTree(_data),
            highlightSearchResults:true,
            searchResultBackColor:'#D3AE30',
            multiSelect: true,
            nodeIcon: "glyphicon glyphicon-unchecked",
            collapseIcon: "glyphicon glyphicon-triangle-bottom",
            expandIcon: 'glyphicon glyphicon-triangle-right',
            backColor: '#F9F9F9',
            borderColor: '#C6C6C6',
            selectedIcon: 'glyphicon glyphicon-check',
            levels: 5,

             onNodeSelected:function(event,data){

                  roleCheckParentNode(data.pid,data.id);
                  
            },
            onNodeUnselected:function(event,data){

                roleCheckChildNode(data.pid,data.id);
            } 
            //showCheckbox:true,
        });
        //$('#tree').treeview('search', [0,{ ignoreCase: false, exactMatch: false }]);
    },
    'json')

}

function initUserGroupMenusTreeMul() {

    $.post('/system/menu/list',
    function(_data) {

        $('#userGroupMenus').treeview({

            data: generateTree(_data),
            highlightSearchResults:true,
            searchResultBackColor:'#D3AE30',
            multiSelect: true,
            nodeIcon: "glyphicon glyphicon-unchecked",
            collapseIcon: "glyphicon glyphicon-triangle-bottom",
            expandIcon: 'glyphicon glyphicon-triangle-right',
            backColor: '#F9F9F9',
            borderColor: '#C6C6C6',
            selectedIcon: 'glyphicon glyphicon-check',
            levels: 5,

             onNodeSelected:function(event,data){

                  userGroupCheckParentNode(data.pid,data.id);
            },
            onNodeUnselected:function(event,data){

                userGroupCheckChildNode(data.pid,data.id);
            } 
            //showCheckbox:true,
        });
        //$('#tree').treeview('search', [0,{ ignoreCase: false, exactMatch: false }]);
    },
    'json')

}

function setUserGroupMenuNodesSelected(data) {

    var nodes = $('#userGroupMenus').treeview('getEnabled');
    var userGroupMenus = new Array();

    for (var i in data) {
        userGroupMenus.push(data[i].Id);
    }

    var len = userGroupMenus.length;
    var len2 = 0;

    //console.info(userGroupMenus);
    for (var i in nodes) {
        if (len2 == len) {
            break;
        }
        for (j in userGroupMenus) {
            if (nodes[i].id == userGroupMenus[j]) {

                len2++;
                $('#userGroupMenus').treeview('selectNode', [nodes[i], {
                    silent: true
                }]);
                break;
            }
        }
    }

    return;
}

function setRoleMenuNodesSelected(data) {

    var nodes = $('#roleMenus').treeview('getEnabled');
    var roleMenus = new Array();

    for (i in data) {
        roleMenus.push(data[i].Id);
    }

    var len = roleMenus.length;
    var len2 = 0;

    for (i in nodes) {
        if (len2 == len) {
            break;
        }
        for (j in roleMenus) {
            if (nodes[i].id == roleMenus[j]) {

                len2++;
                $('#roleMenus').treeview('selectNode', [nodes[i], {
                    silent: true
                }]);
                break;
            }
        }
    }

    return;
}

function setOperMenuNodeSelect(id) {

    var nodes = $('#tree').treeview('getEnabled');
    
    for (var i in nodes) {

        /*if (nodes[i].pid == '-1'){
            $('#tree').treeview('disableNode', [nodes[i], {silent: true}]);
        }*/
        if (id == nodes[i].id) {

            $('#tree').treeview('selectNode', [nodes[i], {
                silent: true
            }]);
        }
    }


    return;
}

function setParentNodeSelected(pids,selectIds) {

    var nodes = $('#my-parentMenus').treeview('getEnabled');
    //console.info(nodes)

    //var v1 ,v2 = false;
    for (var i in nodes) {
       /* for (var j in pids){
            if (pids[j] == nodes[i].id) {

            $('#my-parentMenus').treeview('selectNode', [nodes[i], {silent: true}]);
            break;
            }
        }*/
        for (var j in selectIds){
            if (selectIds[j] == nodes[i].id){

                 $('#my-parentMenus').treeview('disableNode', [nodes[i], {silent: true}]);
                 break;
            }
        }
    }
    return;
}

function checkIsParentContactChild(mainId, selectId){
    var nodes = $('#my-parentMenus').treeview('getEnabled');

    var pN = '';
    for (var i in nodes){
      if (nodes[i].id == selectId){
        pN = $('#my-parentMenus').treeview('getParent', nodes[i].nodeId);
      }
    }


    //console.info(pN);

    if (pN){
      if (pN.pid == mainId){
        return 1
      }
    }
    return 0
}

function getParentMenuSelectIds(){
    var nodes =$('#my-parentMenus').treeview('getSelected');

    var pids = new Array();
    for (var i in nodes){
        pids.push(nodes[i].id);
    }

    return pids;
}


function getSelectNode() {
    return $('#tree').treeview('getSelected');
}

function getRoleSelectNodeIds() {

    var nodes = $('#roleMenus').treeview('getSelected');
    var ids = '';
    for (i in nodes) {
        ids = ids + nodes[i].id + ',';
    }

    if (ids != '') {
        ids = ids.substring(0, ids.length - 1);
    }

    return ids;
}

function getUserGroupSelectNodeIds() {

    var nodes = $('#userGroupMenus').treeview('getSelected');
    var ids = '';
    for (var i in nodes) {
        ids = ids + nodes[i].id + ',';
    }

    if (ids != '') {
        ids = ids.substring(0, ids.length - 1);
    }

    return ids;
}

function userGroupCheckParentNode(pid, nid) {

    var tempPid = pid;
    var nodes = $('#userGroupMenus').treeview('getEnabled');


     for (var i in nodes){
        if (nodes[i].id == pid){

            pid = nodes[i].pid;
            nodes[i].state.selected = true;
            break;
        }
     }

     if (tempPid != pid){
        userGroupCheckParentNode(pid);
     }
    

    return;
}

function userGroupCheckChildNode(pid, nid){

    var nodes = $('#userGroupMenus').treeview('getEnabled');

    for (var i in nodes){
        if (nodes[i].pid == nid){

            nodes[i].state.selected = false;
            userGroupCheckChildNode('',nodes[i].id);
        }
    }
    

    return;
}

function roleCheckParentNode(pid, nid) {

    var tempPid = pid;
    var nodes = $('#roleMenus').treeview('getEnabled');


    for (var i in nodes){
        if (nodes[i].id == pid){

            pid = nodes[i].pid;
            nodes[i].state.selected = true;
            break;
        }
    }

     if (tempPid != pid){
        roleCheckParentNode(pid);
     }
    

    return;
}

function roleCheckChildNode(pid, nid) {

    var nodes = $('#roleMenus').treeview('getEnabled');

    for (var i in nodes){
        if (nodes[i].pid == nid){

            nodes[i].state.selected = false;
            roleCheckChildNode('',nodes[i].id);
        }
    }
    

    return;
}

function generateTree(_data) {

    var tree = new Array();
   // console.info(_data);
    for (var i in _data) {

        if (_data[i].ParentId == '-1') {

            var o = new MenuTree(_data[i].Name, new Array(), _data[i].Id, _data[i].ParentId, _data[i].Icon, new State(false, false, true, false),_data[i].ParentId);

            findChildNodes(_data, o, o.id);
            tree.push(o);
        }
    }

    //console.info(JSON.stringify(tree));
    return JSON.stringify(tree);
}

function findChildNodes(_data, o, pid) {

    var childNodes = new Array();
    for (i in _data) {

        if (_data[i].ParentId == pid) {

            var child = new MenuTree(_data[i].Name, new Array(), _data[i].Id, _data[i].ParentId, _data[i].Icon, new State(false, false, true, false),_data[i].ParentId);

            var smNode = findChildNodes(_data, child, child.id);
            childNodes.push(smNode);
        }
    }

    o.nodes = childNodes;
    return o;
}

function MenuTree(text, nodes, id, pid, icon, state,parentId) {
    var o = new Object();
    o.text = text;
    o.nodes = nodes;
    o.parentId=parentId;
    o.pid = pid;
    o.icon = icon;
    o.state = state;
    o.id = id;
    return o;
}

function State(checked, disabled, expanded, selected) {
    var o = new Object();
    o.checked = checked;
    o.disabled = disabled;
    o.expanded = expanded;
    o.selected = selected;

    return o;
}