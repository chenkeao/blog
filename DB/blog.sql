-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- 主机： localhost:3306
-- 生成日期： 2020-04-26 23:28:59
-- 服务器版本： 10.3.22-MariaDB-1ubuntu1
-- PHP 版本： 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `blog`
--

-- --------------------------------------------------------

--
-- 表的结构 `banner_images`
--

CREATE TABLE `banner_images` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `banner_images`
--

INSERT INTO `banner_images` (`id`, `created_at`, `updated_at`, `path`, `description`, `url`) VALUES
(1, '2020-04-26 23:24:13', '2020-04-26 23:24:13', 'static\\banner\\1587914653IMG_9672.JPG', '前端代码', 'https://cn.bing.com');

-- --------------------------------------------------------

--
-- 表的结构 `comments`
--

CREATE TABLE `comments` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `post_id` int(10) UNSIGNED DEFAULT NULL,
  `read_state` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `game_boards`
--

CREATE TABLE `game_boards` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `score` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `links`
--

CREATE TABLE `links` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `sort` int(11) DEFAULT 0,
  `view` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `pages`
--

CREATE TABLE `pages` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `body` varchar(255) DEFAULT NULL,
  `view` int(11) DEFAULT NULL,
  `is_published` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `posts`
--

CREATE TABLE `posts` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `title` text DEFAULT NULL,
  `body` text DEFAULT NULL,
  `view` int(11) DEFAULT NULL,
  `is_published` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `posts`
--

INSERT INTO `posts` (`id`, `created_at`, `updated_at`, `title`, `body`, `view`, `is_published`) VALUES
(1, '2020-04-26 23:26:02', '2020-04-26 23:26:06', '使用 Hugo + Firebase 搭建博客，以及踩坑', '长久以来学习编程遇到问题都是直接 `Google + CSDN + 简书 + 掘金` 一把梭。时间久了发现很多文章都是从一些个人博客转载过来的，有时偶然点进某些大佬的博客，就仿佛又发现了新大陆一样，不禁感叹大佬就是大佬，不但内容写得好，连网站都做得很漂亮。\r\n\r\n羡慕之余想到自己之前有过一些建站的经验，两年前曾经在腾讯云租过服务器，但是当时真的是菜鸟中的菜鸟，一个服务器就搭了一个只有几个静态页面的小网站，现在想起来真是痛心疾首，不过好在自己也算是进步了不少，正好在学习 Golang 的时候看到 [@ISLAND](https://youngxhui.top/) 这个博客的作者发表的建站教程，就花了半天功夫研究了一下，跟着做了一个出来，下面记录一下过程和遇到的问题。\r\n\r\n## 1. 安装 Hugo\r\n\r\n首先要安装 [Hugo](https://www.gohugo.org/)，[Hugo](https://www.gohugo.org/)是一个由 Golang 编写的静态网站生成器，出于对 Golang 百分之二百的好感，看到后直接就决定选择它了。安装雨果不需要任何依赖，直接在 [Github](https://github.com/gohugoio/hugo/releases) 页面下载对应的版本，安装即可。如果你使用 Chocolatey 的话就更简单了：\r\n\r\n```powershell\r\ncinst hugo -y\r\n```\r\n\r\n## 2. 创建网站\r\n\r\n-   [Hugo](https://www.gohugo.org/)的使用特别简单，安装后直接使用命令：\r\n\r\n```powershell\r\nhugo new site [site_name]\r\n```\r\n\r\n-   然后 [Hugo](https://www.gohugo.org/) 就会光速生成一个站点目录了，接下来进入目录，然后新建一个文章\r\n\r\n```powershell\r\ncd [site_name]\r\nhugo new posts/我的第一篇文章.md\r\n```\r\n\r\n-   [Hugo](https://www.gohugo.org/) 会自动在 `content/posts/`  目录下生成指定的文件，接下来在本地把站点运行起来\r\n\r\n```powershell\r\nhugo server\r\n```\r\n\r\n最后在地址栏输入 `localhost:1313` 就可以看到页面的预览了。如果不出意外的话，现在页面上还是一片空白，但是只要可以访问就说明我们的站点是部署成功的，接下来只要修改一下主题就可以正常显示了。\r\n\r\n## 3. 选择主题\r\n\r\n-   在 [皮肤列表](https://www.gohugo.org/theme/) 选择一个好看的主题，`下载` 或者 ` git clone` 到 `themes` 目录下，再在运行时选择想要的皮肤就好了：\r\n\r\n```powershell\r\nhugo server --theme=hyde\r\n```\r\n\r\n这个时候再打开 `localhost:1313` 就可以看到效果了。:grinning:\r\n\r\n此时我们刚刚创建的文章还没有显示出来，这是因为新建的文章默认都是草稿状态，只需要将文章开头的 `draft: true` 改为 `false` 或者直接删除这个字段就可以了。再回到浏览器刷新下，就可以看到文章已经显示出来了。 \r\n\r\n当然了，如果想要更换为这个主题，则需要再 `config` 文件里进行配置。具体可以参加 Hugo 的 [中文文档](https://www.gohugo.org/)。\r\n\r\n## 4. 托管到 Firebase \r\n\r\n**此过程需要科学上网，如果不能科学上网的话可以使用其他方法**\r\n\r\n一开始选择将页面放到 github 上，但是速度访问实在是太慢了，而且绑定个人的域名和 SSL 比较麻烦，之后选择了 firebase，绑定域名很方便，而且自带 SSL，不需要折腾了。**如果没有 google 账号的话最好先注册一个。**\r\n\r\n-   在 firebase 官网新建一个项目，过程很方便，一直下一步就好了。\r\n\r\n-   在本地安装 [node.js](https://nodejs.org/zh-cn/) 环境，如果使用 Chocolatey：`choco install nodejs`\r\n-   使用 `npm install -g firebase-tools` 命令安装 Firebase CLI 。\r\n-   安装完成后使用 `firebase login` 命令登录\r\n-   在站点的根目录执行 `firebase init` ，接下来会有几个选项：\r\n    -   Are you ready to proceed?(Y/n),这里直接输入 `y`，然后回车\r\n    -   接下来会跳出选项选择项目的类型，使用上下箭头选择 Hosting: Configure and deploy Firebase Hosting sites，然后按空格键确定，回车进入下一个选项。\r\n    -   下面选择 Use an existing project ，使用一个已经存在的项目。\r\n    -   然后应该就可以看到你刚刚创建的项目了，选择之后会询问是否使用 public 文件夹，这里直接回车\r\n    -   最后会询问 Configure as a single-page app (rewrite all urls to /index.html)? (y/N) 这里输入 `n`，然后回车。\r\n-   在站点的根目录执行 `hugo` 命令构建站点。\r\n-   构建成功后使用 `firebase deploy` 命令将文件提交，成功后会给出网站的网站，在浏览器访问就可以看到效果了。\r\n\r\n## 5. 踩坑\r\n\r\n### 1. firebase 无法登录\r\n\r\n使用 shadowsocks 或者 shadowsocksR 的小伙伴可能会遇到 firebase 无法登陆的问题，解决这个问题的方法很简单，就是使用全局代理。   \r\n\r\n但是  ss 或者 ssr 本身是不支持全局代理的，怎么办呢？有一个办法是在路由器上科学上网，这样所有的链接就都是经过代理的了，但是像我这种，使用的路由器不支持科学上的小伙伴咋办呢？在网上找了很久，终于找到了一个软件：`SSTap`。   \r\n\r\n这个软件可以模拟一个网卡，以这种方式实现了全局代理。\r\n\r\n下载链接：[百度网盘](https://pan.baidu.com/s/1BUr-PoJgdSLg3w6y5M3XwQ)  提取码：`7gQO`\r\n\r\n软件的使用很简单，可以直接将 ss 链接导入进去，也可以根据订阅自动获取，代理模式选一下绕过中国大陆的IP，然后记得暂时把 ss 设置为直连模式，再试一下就可以了。', 0, 1),
(2, '2020-04-26 23:26:26', '2020-04-26 23:27:19', '\"科普向 ：人类第一张黑洞照片', '2019年4月10日，人类第一张黑洞照片发布了。\r\n其实黑洞这种天体我们早就耳熟能详，在科幻题材的电影和文学作品里，黑洞都占据着重要的地位。\r\n\r\n## 为什么人类才获得黑洞的照片？\r\n\r\n想搞明白这个问题我们先要知道几个知识：\r\n\r\n1. 万有引力：任何两个物体都是互相吸引的。例如：\r\n   * 地球的引力让我们难以离开地球\r\n   * 月球的引力让海洋形成潮汐\r\n2. 逃逸速度：如果一个物体想离开一个星体，就一定要达到一定的初速度。\r\n   我们想飞离地球就一定要达到一定的初速度(第二宇宙速度：11.2公里/秒)，不然无论如何也会被拉回地面。\r\n3. 根据相对论，有质量物体的速度永远无法达到光速。\r\n\r\n而黑洞因为巨大的引力，逃逸速度超过了光速，也就导致连光也无法逃离黑洞。于是我们就永远无法观测到黑洞的内部。\r\n这里不得不提到一个概念：\r\n\r\n>**事件视界**：因为永远无法观测到黑洞的内部，区分黑洞内外的界限就是“事件视界”。一旦当星体坠入视界之内，就永远无法逃离。因为任何物体都不能达到光速，于是黑洞里也无法传出任何信息。\r\n\r\n就因为这个原因，我们猜测了几百年的黑洞一只得不到应验。\r\n但是，虽然我们无法观测到“视界”内部，但黑洞吸引周围的星体时，会产生大量的X辐射(就是行李安检时用的光)。在X光的波段拍摄，我们才有了第一张黑洞的照片。\r\n\r\n## 黑洞是怎么形成的？\r\n\r\n在理解这个问题之前还需要知道一个概念：\r\n\r\n* 史瓦西半径：一个球体的半径小于它自身的史瓦西半径时，就会无限压缩成一个没有纬度的点，也就是黑洞。\r\n\r\n当一个恒星油尽灯枯，自身无法支撑自己的引力的时候，就会塌陷。要是塌陷到了自己的史瓦西半径之内.....黑洞就产生了。  \r\n太阳会不会变成黑洞？  \r\n几十亿年后，太阳也会有油尽灯枯的一天。但是太阳的质量太小了，远不足以变成黑洞，最后可能的是塌陷成一颗白矮星。质量更大的恒星会塌陷成密度更大的中子星。而质量更大的恒星，就可能会形成黑洞了。\r\n\r\n然而斗转星移，宇宙的尺度太大了。我们的人类文明在时间的长河中或许只是沧海一粟。  \r\n我第一次看到黑洞照片的时候差一点热泪盈眶。它的神秘包容着我从小到大对星空的向往和好奇，遗憾看到的图片都是渲染图或者想象图。   \r\n现在的一张照片或许微不足道，但未来会带来什么，我们或许都想象不到。   ', 1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `post_tags`
--

CREATE TABLE `post_tags` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `post_id` int(10) UNSIGNED DEFAULT NULL,
  `tag_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `smms_files`
--

CREATE TABLE `smms_files` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `store_name` varchar(255) DEFAULT NULL,
  `size` int(11) DEFAULT NULL,
  `width` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  `hash` varchar(255) DEFAULT NULL,
  `delete` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `subscribers`
--

CREATE TABLE `subscribers` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `verify_state` tinyint(1) DEFAULT 0,
  `subscribe_state` tinyint(1) DEFAULT 1,
  `out_time` datetime DEFAULT '2100-12-12 00:00:00',
  `secret_key` varchar(255) DEFAULT NULL,
  `signature` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `subscribers`
--

INSERT INTO `subscribers` (`id`, `created_at`, `updated_at`, `deleted_at`, `email`, `verify_state`, `subscribe_state`, `out_time`, `secret_key`, `signature`) VALUES
(1, '2020-04-26 23:27:39', '2020-04-26 23:28:01', NULL, 'a15961131877@gmail.com', 1, 1, '2020-04-26 23:28:01', '9981b78a-9465-47d7-b5c2-a5a4b71d7d0f', '5267b92b738e9112a1c7e00477c8f65e');

-- --------------------------------------------------------

--
-- 表的结构 `tags`
--

CREATE TABLE `tags` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `telephone` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `verify_state` varchar(255) DEFAULT '0',
  `secret_key` varchar(255) DEFAULT NULL,
  `out_time` datetime DEFAULT '2100-12-12 00:00:00',
  `github_login_id` varchar(255) DEFAULT NULL,
  `github_url` varchar(255) DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT NULL,
  `avatar_url` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `lock_state` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `email`, `telephone`, `password`, `verify_state`, `secret_key`, `out_time`, `github_login_id`, `github_url`, `is_admin`, `avatar_url`, `nick_name`, `lock_state`) VALUES
(1, '2020-04-26 23:22:43', '2020-04-26 23:22:43', NULL, 'admin@qq.com', NULL, '5c2c4e44e67d866e1122013bb75bc66e', '0', NULL, '2100-12-12 00:00:00', NULL, '', 1, '', '', 0);

--
-- 转储表的索引
--

--
-- 表的索引 `banner_images`
--
ALTER TABLE `banner_images`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `game_boards`
--
ALTER TABLE `game_boards`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `links`
--
ALTER TABLE `links`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_links_deleted_at` (`deleted_at`);

--
-- 表的索引 `pages`
--
ALTER TABLE `pages`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `post_tags`
--
ALTER TABLE `post_tags`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uk_post_tag` (`post_id`,`tag_id`);

--
-- 表的索引 `smms_files`
--
ALTER TABLE `smms_files`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `subscribers`
--
ALTER TABLE `subscribers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uix_subscribers_email` (`email`),
  ADD KEY `idx_subscribers_deleted_at` (`deleted_at`);

--
-- 表的索引 `tags`
--
ALTER TABLE `tags`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uix_users_email` (`email`),
  ADD UNIQUE KEY `uix_users_telephone` (`telephone`),
  ADD UNIQUE KEY `uix_users_github_login_id` (`github_login_id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `banner_images`
--
ALTER TABLE `banner_images`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `game_boards`
--
ALTER TABLE `game_boards`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `links`
--
ALTER TABLE `links`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `pages`
--
ALTER TABLE `pages`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `post_tags`
--
ALTER TABLE `post_tags`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `smms_files`
--
ALTER TABLE `smms_files`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `subscribers`
--
ALTER TABLE `subscribers`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `tags`
--
ALTER TABLE `tags`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
