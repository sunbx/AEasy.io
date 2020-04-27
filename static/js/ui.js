// 页面的脚本逻辑
$(function () {



//	============================================================

    $('#go-register').on('click', function () {
        $('.div-login').hide();
        $('.div-register').show();
    });

    $('#go-login').on('click', function () {
        $('.div-register').hide();
        $('.div-login').show();
    });
    $.LoadingOverlaySetup({
        image: "../static/img/android-chrome-256x256.png",
        size: "15",
        background: "#ffffff",
        color: "#ba0c55"
    });
//	============================================================
    let isCountDown = false;
    let goCaptcha = $('#go-captcha').on('click', function () {

        if (!isCountDown) {
            let email = $('#register-input-email').val();
            if (email === '') {
                swal('Email is null');
                return;
            }
            $.LoadingOverlay("show");

            setTimeout(function () {
                $.post("/mail/send", {email: email, type: "1"}, function (result) {
                    let obj = jQuery.parseJSON(result);
                    if (obj.code === 200) {
                        countDown(60);
                    } else {
                        swal(obj.msg);
                    }
                    $.LoadingOverlay("hide");
                });
            }, 1500);


        }

        //读秒
        function countDown(count) {
            isCountDown = true;
            const down = setInterval(CountDown, 1000);//每秒执行一次，赋值
            function CountDown() {

                goCaptcha.text(count + 's');//写入
                goCaptcha.css("background-image", "linear-gradient(116deg, #DCDCDC 0%, #DCDCDC 10%)");
                if (count === 0) {
                    $('#go-captcha').text('Get');//
                    $("#go-captcha").css("background-image", "linear-gradient(116deg, #5E3AFF 0%, #F7296E 90%)");
                    // $('#putvf1').css('display', 'none');//修改状态
                    // $('#putvf').css('display', 'block');
                    clearInterval(down);//销毁计时器
                    isCountDown = false;
                    return;
                }
                count--;
            }
        }
    });

    $('#register').on('click', function () {
        let email = $('#register-input-email').val();
        let captcha = $('#register-input-captcha').val();
        let password = $('#register-input-password').val();

        if (!isPasswd(password)) {
            swal('Only 6-20 letters, Numbers, and underscores can be entered');
            return;
        }
        $.post("/user/register", {
            email: email,
            captcha: captcha,
            password: $.md5(password + 'aeasy')
        }, function (result) {
            let obj = jQuery.parseJSON(result);
            if (obj.code === 200) {
                window.location.href = '/';
            } else {
                swal(obj.msg);
            }
        });
    });

    $('#login').on('click', function () {
        let email = $('#login-input-email').val();
        let password = $('#login-input-password').val();

        if (isPasswd(password)) {
            $.post("/user/login", {email: email, password: $.md5(password + 'aeasy')}, function (result) {
                let obj = jQuery.parseJSON(result);
                if (obj.code === 200) {
                    window.location.href = '/';
                } else {
                    swal(obj.msg);
                }
            });
        } else {
            swal('Only 6-20 letters, Numbers, and underscores can be entered');
        }
    });


    $('#go-forget').on('click', function () {
        swal('Please contact customer service to retrieve the password')
    });
    $('#go-privacy').on('click', function () {
        var options = $("#select option:selected"); //获取选中的项

        alert(options.val()); //拿到选中项的值

        alert(options.text()); //拿到选中项的文本

        // swal('This website is used for development and learning, if the abnormal situation is not responsible, the final right of interpretation belongs to this website')
    });
    $('#go-logout').on('click', function () {
        $.post("/user/logout", function (result) {
            let obj = jQuery.parseJSON(result);
            if (obj.code === 200) {
                window.location.href = '/';
            } else {
                swal(obj.msg);
            }
        });
    });
    let address = $('#address').text();
    $('#qrcode').qrcode({
        foreground: "#311B58",
        background: "#FFF",
        width: 200, height: 200,
        text: address
    });

});

//校验密码：只能输入6-20个字母、数字、下划线
function isPasswd(s) {
    let patrn = /^(\w){6,20}$/;
    return patrn.exec(s);

}

