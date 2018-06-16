function initUserGroupOperTabs(userGroupIds) {

    $('#userGroupOpers-ul').empty();
    $('#userGroupOpers-content').empty();

    if (userGroupIds == '') {
        return;
    }

    $.post('/system/userGroup/userGroupToContactOper', {
        userGroupId: userGroupIds
    },
    function(_data) {

        //console.info(_data);
        _data = $.parseJSON(_data);
        var menuUnitOper = _data["menuUnitOper"];
        var userGroupOper = _data["userGroupOpers"];
        var menuNames = _data["menuNames"];

        //console.info(menuUnitOper);console.info(userGroupOper);console.info(menuNames);
        //console.info(userGroupOper);
        var menuOpers = '';

        var ul = $('#userGroupOpers-ul');
        var content = $('#userGroupOpers-content');
        var lis = '';
        var dropdown = '';
        var contents = '';
        var opers = '';
        var tag = 0;

        for (var i in menuUnitOper) {

            tag++;
            if (tag > 5) {
                if (tag == 6) {
                    dropdown = dropdown + '<li class="dropdown"><a href="#" id="myTabDrop1" class="dropdown-toggle"' + 'data-toggle="dropdown">鏇村�<b class="caret"></b> </a><ul class="dropdown-menu" role="menu" aria-labelledby="myTabDrop1">' + '<li><a href="#' + i + '" tabindex="-1" data-toggle="tab">' + menuNames[i] + '</a></li>';
                } else {
                    dropdown = dropdown + '<li><a href="#' + i + '" tabindex="-1" data-toggle="tab">' + menuNames[i] + '</a></li>';
                }

            } else {

                if (tag == 1) {

                    lis = lis + '<li class="active"><a href="#' + i + '" data-toggle="tab">' + menuNames[i] + '</a></li>';
                } else {

                    lis = lis + '<li><a href="#' + i + '" data-toggle="tab">' + menuNames[i] + '</a></li>';
                }

            }

            opers = '<div>';
            menuOpers = menuUnitOper[i];
            for (var j in menuOpers) {

                if (j % 5 == 0) {
                    opers = opers + '<br><br>';
                }

                if (checkOperId(userGroupOper, menuOpers[j].Id)){

                    opers = opers + '<label class="checkbox-inline"><input type="checkbox" checked name="operCheck" value="' + menuOpers[j].Id + '">' + menuOpers[j].Name + '</label>';
                }else{

                    opers = opers + '<label class="checkbox-inline"><input type="checkbox" name="operCheck" value="' + menuOpers[j].Id + '">' + menuOpers[j].Name + '</label>';
                }
                
            }
            opers = opers + '</div>';

            if (tag == 1) {

                contents = contents + '<div class="tab-pane fade in active" id="' + i + '">' + opers + '</div>';
            } else {

                contents = contents + '<div class="tab-pane fade" id="' + i + '">' + opers + '</div>';

            }

        }

        dropdown = dropdown + '</ul></li>';

        lis = lis + dropdown;

        //console.info(lis);
        //console.info(contents);
        ul.append(lis);
        content.append(contents);
        $('#userGroupOpers').append(ul).append(content);
        $('#userGroupOperModal').modal('show', true);
    })
}


function checkOperId(userGroupS , operId){

    for (var i in userGroupS){

        if (userGroupS[i].Id == operId){
            return true;
        }
    }
    return false;
}

function isNotEmptyObject(obj) {

    for (var i in obj) {
        return true;
    }

    return false;
}




