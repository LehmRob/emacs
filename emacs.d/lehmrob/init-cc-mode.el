(add-to-list 'auto-mode-alist '("\\.h\\'" . c++-mode))
(c-add-style "lehmrob-style"
			 '("stroustrup"
			   (c-basic-offset . 2)
			   (c-set-offset 'substatement-open '0)
			   (c-set-offset 'case-label '+)
			   (c-offsets-alist . ((innamespace . [0])))))

(defun my-cpp-mode-hook()
  (c-set-style "lehmrob-style")
  (auto-fill-mode)
  (c-toggle-hungry-state 1))

(add-hook 'c++-mode-hook 'my-cpp-mode-hook)

(require 'clang-format)
(global-set-key (kbd "C-c i") 'clang-format-region)
(global-set-key (kbd "C-c u") 'clang-format-buffer)


(provide 'init-cc-mode)
