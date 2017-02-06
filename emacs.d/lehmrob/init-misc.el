(setq inhibit-startup-message t)
(fset 'yes-or-no-p 'y-or-n-p)
(toggle-scroll-bar -1)
(menu-bar-mode -99)
(tool-bar-mode -1)
(setq visible-bell t)
;; "C-h d transient" for more info
(setq transient-mark-mode t)
(pending-delete-mode t)
(display-time-mode 1)
(emacs-lock-mode)
(column-number-mode)
(set-default 'indicate-empty-lines t)
(setq-default default-tab-width 4)
(setq indent-line-function 'insert-tab)
(setq-default display-time-24hr-format t)
;; i want a mouse yank to be inserted where the point is, not where i click
(setq-default mouse-yank-at-point t)
(when (fboundp 'electric-indent-mode) (electric-indent-mode -1))
(defvar whitespace-cleanup-on-save t)

;; Add melpa repository
(require 'package)
(add-to-list 'package-archives
             '("melpa" . "http://melpa.org/packages/") t)
(when (< emacs-major-version 24)
  ;; For important compatibility libraries like cl-lib
  (add-to-list 'package-archives '("gnu" . "http://elpa.gnu.org/packages/"))
  (add-to-list 'package-archives '("marmalade" . "http://marmalade-repo.org/packages/")))
(package-initialize)

(provide 'init-misc)
