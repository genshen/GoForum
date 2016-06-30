/**
 * Created by 根深 on 2016/6/21.
 */
$.forms = {
    "options": {
        form: null,
        submitBtn: null,
        fields: []
    },
    "addValid": function (e) {
        var field = $(e.field);
        field.on("focus", function () {
            valid(field, e.rules);
        }).on("blur", function () {
            valid(field, e.rules);
        }).on("input", function () {  //在输入框内容变化的时候不会触发change，当鼠标在其他地方点一下才会触发
            valid(field, e.rules);
        });
    },
    "init": function (options) {
        this.options = $.extend({}, this.options, options);
        this.options.fields.forEach(function (e) {
            $.forms.addValid(e);
        });

        if (this.options.submitBtn != null) {
            $(this.options.submitBtn).on("click", function () {
                if($($.forms.options.submitBtn).attr("disabled") == "disabled"){
                    return;
                }
                var pass = true;
                $.forms.options.fields.forEach(function (e) {
                    var field = $(e.field);
                    var t = valid(field, e.rules);
                    if (!t && pass) {
                        pass = t;
                        field.focus();
                    }
                });
                if (pass) {
                    var form = $($.forms.options.form);
                    if (form) {
                        form.submit();
                    }
                }
            });
        }
    }
};


function valid(input, rules) {
    var rule = rules.required;
    var value = input.val();

    if (rule && !validRequired(value, input, rule, false)) {
        return false;
    }
    rule = rules.min;
    if (rule && !validMin(value, input, rule)) {
        return false;
    }
    rule = rules.max;
    if (rule && !validMax(value, input, rule)) {
        return false;
    }
    rule = rules.email;
    if (rule && !validEmail(value, input, rule)) {
        return false;
    }
    rule = rules.number;
    if (rules && !validNumber(value, input, rule)) {
        return false;
    }
    rule = rules.regexp;
    if (rule && !validRegexp(value, input, rule)) {
        return false;
    }

    rule = rules.same;
    if (rule && !validSame(value, input, rule)) {
        return false;
    }

    rule = rules.customize;
    if (rule && !validCustomize(value, input, rule)) {
        return false;
    }
    clearError(input);
    return true;
}

function clearError(input) {
    var $formGroup = input.closest(".form-group");
    $formGroup.removeClass("has-error");
    $formGroup.find(".help-block").remove();
}

function setError(input, message) {
    var $formGroup = input.closest(".form-group");
    $formGroup.addClass("has-error");
    var help = $formGroup.find(".help-block");
    if (help.length === 0) {
        input.after("<p class='help-block'>" + message + "</p>");
    } else {
        help.text(message);
    }
}

function validRequired(value, input, rule, isSubmit) {
    if (value) {
        return true;
    } else if (isSubmit) {
        setError(input, rule.message);
        return false;
    } else {
        if (!rule.submitOnly) {
            setError(input, rule.message);
            return false;
        }
    }
    return true;
}
function validMin(value, input, rule) {
    if (value.length >= rule.count) {
        return true;
    }
    setError(input, rule.message);
    return false;
}

function validMax(value, input, rule) {
    if (value.length > rule.count) {
        setError(input, rule.message);
        return false;
    }
    return true;
}

function validEmail(value, input, rule) {
    var Regex = /^(?:\w+\.?)*\w+@(?:\w+\.)*\w+$/;
    if (Regex.test(value)) {
        return true;
    }
    setError(input, rule.message);
    return false;
}

function validNumber(value, input, rule) {
    return true;
}

function validRegexp(value,input, rule) {
    return true;
}

/* rule.element must be a input ,
and we just set one input with error when they are not the same */
function validSame(value,input, rule) {
    var ele = $(rule.element);
    console.log(ele.val()+":=:"+value);
    if(ele.val() == value){
        clearError(ele);
        return true;
    }
    setError(input, rule.message);
    // setError(ele, rule.message);
    return false;
}

function validCustomize(value,input, rule) {
    if(rule.func(value,input)){
        return true;
    }
    setError(input, rule.message);
    return false;
}

//tools
function base64_decode (encodedData) {
    if (typeof window !== 'undefined') {
        if (typeof window.atob !== 'undefined') {
            return decodeURIComponent(unescape(window.atob(encodedData)))
        }
    } else {
        return new Buffer(encodedData, 'base64').toString('utf-8')
    }

    var b64 = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
    var o1,o2,o3,h1,h2,h3,h4,bits;
    var i = 0;
    var ac = 0;
    var dec = '';
    var tmpArr = [];

    if (!encodedData) {
        return encodedData
    }

    encodedData += '';

    do {
        // unpack four hexets into three octets using index points in b64
        h1 = b64.indexOf(encodedData.charAt(i++));
        h2 = b64.indexOf(encodedData.charAt(i++));
        h3 = b64.indexOf(encodedData.charAt(i++));
        h4 = b64.indexOf(encodedData.charAt(i++));

        bits = h1 << 18 | h2 << 12 | h3 << 6 | h4;

        o1 = bits >> 16 & 0xff;
        o2 = bits >> 8 & 0xff;
        o3 = bits & 0xff;

        if (h3 === 64) {
            tmpArr[ac++] = String.fromCharCode(o1)
        } else if (h4 === 64) {
            tmpArr[ac++] = String.fromCharCode(o1, o2)
        } else {
            tmpArr[ac++] = String.fromCharCode(o1, o2, o3)
        }
    } while (i < encodedData.length);

    dec = tmpArr.join('');
    return decodeURIComponent(escape(dec.replace(/\0+$/, '')))
}