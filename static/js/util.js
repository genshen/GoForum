/*util：
 * including:
 * 1:material design tab.js
 * 2.form-validate.js
 * 3.base64_decode
 * 4.jquery.cookie.js
 * 5.snackbar.js
 * */
var Config = {
    home:{title:"Clothes Plus",subtitle:"副标题",logo:"/static/img/logo.png",cover:"/static/img/top_cover.jpg"},
    listMax: {hot_post: 20, post_list: 20, comment: 20}
};
// tab switch
(function ($) {
    'use strict';
    $.fn.tabSwitch = function (oldTab) {
        var $this = $(this),
            $thisNav = $this.closest('.tab-nav'),
            $thisNavIndicator = $('.tab-nav-indicator', $thisNav),
            thisLeft = $this.offset().left,
            thisNavLeft = $thisNav.offset().left,
            thisNavWidth = $thisNav.outerWidth();

        if (oldTab !== undefined && oldTab[0] !== undefined) {
            var oldTabLeft = oldTab.offset().left;
            $thisNavIndicator.css({
                left: (oldTabLeft - thisNavLeft),
                right: (thisNavLeft + thisNavWidth - oldTabLeft - oldTab.outerWidth())
            });

            if (oldTab.offset().left > thisLeft) {
                $thisNavIndicator.addClass('reverse');
                $thisNavIndicator.one('webkitTransitionEnd oTransitionEnd msTransitionEnd transitionend', function () {
                    $thisNavIndicator.removeClass('reverse');
                });
            }
            $thisNavIndicator.addClass('animate').css({
                left: (thisLeft - thisNavLeft),
                right: (thisNavLeft + thisNavWidth - thisLeft - $this.outerWidth())
            }).one('webkitTransitionEnd oTransitionEnd msTransitionEnd transitionend', function () {
                $thisNavIndicator.removeClass('animate');
            });
        }


        return this;
    }
})(jQuery);

$(function () {
    'use strict';

    $('.tab-nav').each(function () {
        $(this).append('<div class="tab-nav-indicator"></div>');
    });

    $(document).on('show.bs.tab', '.tab-nav a[data-toggle="tab"]', function (e) {
        $(e.target).tabSwitch($(e.relatedTarget));
    });
});

//form-validate
/**
 * Created by 根深 on 2016/6/21.
 */
$.forms = {
    "options": {
        form: null,
        submitBtn: null,
        onSubmit: null,
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
            var subBtn = $(this.options.submitBtn);
            subBtn.on("click", function () {
                if (subBtn.attr("disabled") == "disabled") {
                    return;
                }
                var pass = true;
                $.forms.options.fields.forEach(function (e) {
                    var field = $(e.field);
                    var t = valid(field, e.rules, true);
                    if (!t && pass) {
                        pass = t;
                        field.focus();
                    }
                });
                if (pass) {
                    if ($.forms.options.onSubmit) {
                        $.forms.options.onSubmit(subBtn)
                    } else {
                        var form = $($.forms.options.form);
                        if (form) {
                            form.submit();
                        }
                    }
                }
            });
        }
    }
};
function valid(input, rules, isSubmit) {
    var rule = rules.required;
    var value = input.val();
    if (rule && !validRequired(value, input, rule, isSubmit)) {
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
    if ((!isSubmit && rule.submitOnly) || value) {
        return true;
    } else {
        setError(input, rule.message);
        return false;
    }
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

function validRegexp(value, input, rule) {
    return true;
}

/* rule.element must be a input ,
 and we just set one input with error when they are not the same */
function validSame(value, input, rule) {
    var ele = $(rule.element);
    if (ele.val() == value) {
        clearError(ele);
        return true;
    }
    setError(input, rule.message);
    // setError(ele, rule.message);
    return false;
}

function validCustomize(value, input, rule) {
    if (rule.func(value, input)) {
        return true;
    }
    setError(input, rule.message);
    return false;
}

//base64_decode
function base64_decode(encodedData) {
    if (typeof window !== 'undefined') {
        if (typeof window.atob !== 'undefined') {
            return decodeURIComponent(unescape(window.atob(encodedData)))
        }
    } else {
        return new Buffer(encodedData, 'base64').toString('utf-8')
    }

    var b64 = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
    var o1, o2, o3, h1, h2, h3, h4, bits;
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

//js-cookie
/*!
 * JavaScript Cookie v2.1.2
 * https://github.com/js-cookie/js-cookie
 *
 * Copyright 2006, 2015 Klaus Hartl & Fagner Brack
 * Released under the MIT license
 */
;(function (factory) {
    if (typeof define === 'function' && define.amd) {
        define(factory);
    } else if (typeof exports === 'object') {
        module.exports = factory();
    } else {
        var OldCookies = window.Cookies;
        var api = window.Cookies = factory();
        api.noConflict = function () {
            window.Cookies = OldCookies;
            return api;
        };
    }
}(function () {
    function extend() {
        var i = 0;
        var result = {};
        for (; i < arguments.length; i++) {
            var attributes = arguments[i];
            for (var key in attributes) {
                result[key] = attributes[key];
            }
        }
        return result;
    }

    function init(converter) {
        function api(key, value, attributes) {
            var result;
            if (typeof document === 'undefined') {
                return;
            }
            // Write
            if (arguments.length > 1) {
                attributes = extend({
                    path: '/'
                }, api.defaults, attributes);

                if (typeof attributes.expires === 'number') {
                    var expires = new Date();
                    expires.setMilliseconds(expires.getMilliseconds() + attributes.expires * 864e+5);
                    attributes.expires = expires;
                }

                try {
                    result = JSON.stringify(value);
                    if (/^[\{\[]/.test(result)) {
                        value = result;
                    }
                } catch (e) {
                }

                if (!converter.write) {
                    value = encodeURIComponent(String(value))
                        .replace(/%(23|24|26|2B|3A|3C|3E|3D|2F|3F|40|5B|5D|5E|60|7B|7D|7C)/g, decodeURIComponent);
                } else {
                    value = converter.write(value, key);
                }

                key = encodeURIComponent(String(key));
                key = key.replace(/%(23|24|26|2B|5E|60|7C)/g, decodeURIComponent);
                key = key.replace(/[\(\)]/g, escape);

                return (document.cookie = [
                    key, '=', value,
                    attributes.expires && '; expires=' + attributes.expires.toUTCString(), // use expires attribute, max-age is not supported by IE
                    attributes.path && '; path=' + attributes.path,
                    attributes.domain && '; domain=' + attributes.domain,
                    attributes.secure ? '; secure' : ''
                ].join(''));
            }

            // Read

            if (!key) {
                result = {};
            }

            // To prevent the for loop in the first place assign an empty array
            // in case there are no cookies at all. Also prevents odd result when
            // calling "get()"
            var cookies = document.cookie ? document.cookie.split('; ') : [];
            var rdecode = /(%[0-9A-Z]{2})+/g;
            var i = 0;

            for (; i < cookies.length; i++) {
                var parts = cookies[i].split('=');
                var cookie = parts.slice(1).join('=');

                if (cookie.charAt(0) === '"') {
                    cookie = cookie.slice(1, -1);
                }

                try {
                    var name = parts[0].replace(rdecode, decodeURIComponent);
                    cookie = converter.read ?
                        converter.read(cookie, name) : converter(cookie, name) ||
                    cookie.replace(rdecode, decodeURIComponent);

                    if (this.json) {
                        try {
                            cookie = JSON.parse(cookie);
                        } catch (e) {
                        }
                    }

                    if (key === name) {
                        result = cookie;
                        break;
                    }

                    if (!key) {
                        result[name] = cookie;
                    }
                } catch (e) {
                }
            }

            return result;
        }

        api.set = api;
        api.get = function (key) {
            return api(key);
        };
        api.getJSON = function () {
            return api.apply({
                json: true
            }, [].slice.call(arguments));
        };
        api.defaults = {};

        api.remove = function (key, attributes) {
            api(key, '', extend(attributes, {
                expires: -1
            }));
        };

        api.withConverter = init;

        return api;
    }

    return init(function () {
    });
}));

/* SnackbarJS - MIT LICENSE (https://github.com/FezVrasta/snackbarjs/blob/master/LICENSE.md) */
(function (c) {
    function d(a) {
        return "undefined" !== typeof a && null !== a ? !0 : !1
    }

    c(document).ready(function () {
        c("body").append("<div id=snackbar-container/>")
    });
    c(document).on("click", "[data-toggle=snackbar]", function () {
        c(this).snackbar("toggle")
    }).on("click", "#snackbar-container .snackbar", function () {
        c(this).snackbar("hide")
    });
    c.snackbar = function (a) {
        if (d(a) && a === Object(a)) {
            var b;
            b = d(a.id) ? c("#" + a.id) : c("<div/>").attr("id", "snackbar" + Date.now()).attr("class", "snackbar");
            var g = b.hasClass("snackbar-opened");
            d(a.style) ? b.attr("class", "snackbar " + a.style) : b.attr("class", "snackbar");
            a.timeout = d(a.timeout) ? a.timeout : 3E3;
            d(a.content) && (b.find(".snackbar-content").length ? b.find(".snackbar-content").text(a.content) : b.prepend("<span class=snackbar-content>" + a.content + "</span>"));
            d(a.id) ? b.insertAfter("#snackbar-container .snackbar:last-child") : b.appendTo("#snackbar-container");
            d(a.action) && "toggle" == a.action && (a.action = g ? "hide" : "show");
            var e = Date.now();
            b.data("animationId1", e);
            setTimeout(function () {
                b.data("animationId1") === e && (d(a.action) && "show" != a.action ? d(a.action) && "hide" == a.action && b.removeClass("snackbar-opened") : b.addClass("snackbar-opened"))
            }, 50);
            var f = Date.now();
            b.data("animationId2", f);
            0 !== a.timeout && setTimeout(function () {
                b.data("animationId2") === f && b.removeClass("snackbar-opened")
            }, a.timeout);
            return b
        }
        return !1
    };
    c.fn.snackbar = function (a) {
        var b = {};
        if (this.hasClass("snackbar")) {
            b.id = this.attr("id");
            if ("show" === a || "hide" === a || "toggle" == a)b.action = a;
            return c.snackbar(b)
        }
        d(a) && "show" !== a && "hide" !== a && "toggle" != a || (b = {
            content: c(this).attr("data-content"),
            style: c(this).attr("data-style"),
            timeout: c(this).attr("data-timeout")
        });
        d(a) && (b.id = this.attr("data-snackbar-id"), "show" === a || "hide" === a || "toggle" == a) && (b.action = a);
        a = c.snackbar(b);
        this.attr("data-snackbar-id", a.attr("id"));
        return a
    }
})(jQuery);

var Util = {
    parseError: {
        options: {},
        init: function (errors, options) {
            this.options = $.extend({}, { //default options
                snackTimeout: 3000,
                errorCallback: function (message, name) {
                    $.snackbar({content: message, timeout: this.snackTimeout})
                }
            }, options);
            this.execute(errors)
        },
        execute: function (Errors) {
            for (var key in Errors) {
                var err = Errors[key].Errors;
                if (err.length > 0) {
                    this.options.errorCallback(err[0].Message, err[0].Name); //Name == key
                    return false
                }
            }
            return true;
        }
    },
    simpleParseError: {
        options: {
            authUrl: "/account/signin",
            snackTimeout: 3000,
            errorCallback: function (message) {
                $.snackbar({content: message, timeout: this.snackTimeout})
            },
            onUnAuth: function () {
                var url = this.authUrl + "?next=" + document.location.pathname;
                $.snackbar({content: "请<a href='" + url + "'>登录</a>后进行该操作", timeout: this.snackTimeout})
            },
            onSuccess: null
        },
        init: function (status, error, options) {
            this.options = $.extend({}, this.options, options);
            this.execute(status, error)
        },
        execute: function (status, error) {
            switch (status) {
                case 0:
                    this.options.errorCallback(error);
                    return false;
                case 1:
                    if (this.options.onSuccess != null) {
                        this.options.onSuccess();
                    }
                    return true;
                case 3: //unauth
                    this.options.onUnAuth();
                    return false;
            }
        }
    },
    formatTime: function (value) {
        if (typeof value != "number") {
            var v = Date.parse(value);
            if (isNaN(v)) {
                value = (new Date).getTime();
            } else {
                value = v;
            }
        }
        now = (new Date).getTime();
        if (now - value < 60 * 1000) {
            return "刚刚";
        }
        if (now - value < 60 * 60 * 1000) {
            var min = parseInt((now - value) / (60 * 1000));
            return min + "分钟前";
        }
        if (now - value < 24 * 60 * 60 * 1000) {
            var hour = parseInt((now - value) / (60 * 60 * 1000));
            return hour + "小时前";
        }
        if (now - value < 20 * 24 * 60 * 60 * 1000) {
            var day = parseInt((now - value) / (24 * 60 * 60 * 1000));
            return day + "天前";
        }
        var d = new Date(value);
        return d.getFullYear() + "-" + (d.getMonth() + 1) + "-" + d.getDate();
    }
};